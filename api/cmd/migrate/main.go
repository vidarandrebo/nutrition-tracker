package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
)

type Flags struct {
	MigrationDir string
}

func main() {
	opt, err := configuration.ParseOptions("appsettings.json")
	if err != nil {
		panic("read of config failed")
	}

	args := Flags{}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "\tmigrate\n")
		fmt.Fprintf(os.Stderr, "\tdestroy\n")
	}
	flag.StringVar(&args.MigrationDir, "migration-dir", "./", "path to migration directory")
	flag.Parse()

	commands := flag.Args()

	db, err := sql.Open("pgx", opt.DBConnectionString)
	if err != nil {
		panic(err)
	}

	app := DbMigrationApp{
		options: opt,
		flags:   &args,
		db:      db,
	}

	for _, command := range commands {
		switch command {
		case "update":
			app.update()
		case "destroy":
			app.destroy()
		}
	}
}

type TblMigrationHistory struct {
	Id          int64
	Name        string
	DateApplied time.Time
}

type DbMigrationApp struct {
	options *configuration.Options
	flags   *Flags
	db      *sql.DB
}

func (ma *DbMigrationApp) destroy() {
	file, err := os.ReadFile(path.Join(ma.flags.MigrationDir, "destroy.sql"))
	if err != nil {
		panic(err)
	}
	query := string(file)
	ma.db.Exec(query)
	fmt.Println("destroying db")
}
func (ma *DbMigrationApp) update() {
	fmt.Println("migrating to new version of db")

	dirEntries, err := os.ReadDir(ma.flags.MigrationDir)
	if err != nil {
		panic(err)
	}

	validMigrationExpr, err := regexp.Compile("^(([12]\\d{3})-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01])-([a-zA-Z\\-]*).sql)$")

	if err != nil {
		panic(err)
	}

	for _, dirEntry := range dirEntries {
		if validMigrationExpr.MatchString(dirEntry.Name()) {
			fmt.Println(dirEntry.Name())
			file, err := os.ReadFile(path.Join(ma.flags.MigrationDir, dirEntry.Name()))
			if err != nil {
				panic(err)
			}
			query := string(file)
			ma.applyMigration(query, dirEntry.Name())
		}
	}
}

func (ma *DbMigrationApp) applyMigration(query string, name string) bool {
	existing := TblMigrationHistory{}

	getMigErr := ma.db.QueryRow(`
		SELECT id, name, date_applied
		FROM migration_history
		WHERE name = $1
    `, name).Scan(&existing.Id, &existing.Name, &existing.DateApplied)

	if getMigErr == nil {
		fmt.Printf("Migration with id %d and name %s was applied %v\n", existing.Id, existing.Name, existing.DateApplied)
		return false
	}
	fmt.Printf("New migration found applying migration %s\n", name)

	_, runQueryErr := ma.db.Exec(query)
	if runQueryErr != nil {
		fmt.Println(runQueryErr.Error())
		panic(runQueryErr)
	}
	_, insertMigErr := ma.db.Exec(`
		INSERT INTO migration_history(name, date_applied) 
		VALUES ($1, NOW())`, name)
	if insertMigErr != nil {
		panic(insertMigErr)
	}
	return true

	return true

}
