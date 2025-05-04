package internal

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Importer struct {
	Server   http.Server
	DB       *sql.DB
	Options  *configuration.Options
	Logger   *slog.Logger
	Services *Services
	Stores   *Stores
}

func (a *Importer) CloseDB() {
	a.DB.Close()
}

func NewImporter() *Importer {
	return &Importer{}
}
func (a *Importer) configureLogger() {
	logFile, err := os.OpenFile("./dataimporter.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}
	logHandlerOpts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logWriter := io.MultiWriter(logFile, os.Stderr)
	logHandler := slog.NewTextHandler(logWriter, &logHandlerOpts)

	a.Logger = slog.New(logHandler)
}

func (a *Importer) configureDB() {
	db, err := sql.Open("pgx", a.Options.DBConnectionString)
	if err != nil {
		panic(err)
	}
	a.DB = db
}
func (a *Importer) readConfiguration() {
	opt, err := configuration.ParseOptions("appsettings.json")
	if err != nil {
		panic("read of config failed")
	}
	a.Options = opt
}
func (a *Importer) configureServices() {
	a.Services = &Services{}
	a.Services.JwtService = auth.NewJwtService()
	a.Services.HashingService = auth.NewHashingService()
	a.Services.AuthService = auth.NewAuthService(a.Stores.UserStore, a.Services.HashingService)
}
func (a *Importer) configureStores() {
	a.Stores = &Stores{}
	a.Stores.FoodItemStore = fooditem.NewStore(a.DB)
	a.Stores.UserStore = user.NewStore(a.DB, a.Logger)
}

func (a *Importer) Setup() {
	a.readConfiguration()
	a.configureLogger()
	a.configureDB()
	a.configureStores()
	a.configureServices()

}
