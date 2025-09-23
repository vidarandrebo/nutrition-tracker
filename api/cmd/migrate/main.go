package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"regexp"

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

	for _, command := range commands {
		switch command {
		case "update":
			file, err := os.ReadFile("./migrations/2025-09-23-initial.sql")
			fmt.Println("migrating to new version of db")
			if err != nil {
				panic(err)
			}
			dirEntries, err := os.ReadDir(args.MigrationDir)
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
				}
			}
			query := string(file)
			db.Exec(query)
		case "destroy":
			file, err := os.ReadFile("./migrations/destroy.sql")
			if err != nil {
				panic(err)
			}
			query := string(file)
			db.Exec(query)
			fmt.Println("destroying db")
		}
	}
}
