package internal

import (
	"database/sql"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/meal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/middleware"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
)

type Application struct {
	Server      http.Server
	DB          *sql.DB
	Options     *configuration.Options
	Logger      *slog.Logger
	Services    *Services
	Stores      *Stores
	Endpoints   *Endpoints
	Middlewares *Middlewares
}

type Services struct {
	JwtService     *auth.JwtService
	AuthService    *auth.Service
	HashingService *auth.HashingService
}
type Stores struct {
	FoodItemStore *fooditem.Store
	UserStore     *user.Store
	MealStore     *meal.Store
	RecipeStore   *recipe.Store
}

type Endpoints struct {
	FoodItemEndpoint *foodItemEndpoint
	RecipeEndpoint   *recipeEndpoint
	AuthEndpoint     *authEndpoint
	MealEndpoint     *mealEndpoint
}

type Middlewares struct {
	Auth            *middleware.Auth
	RequestMetadata *middleware.RequestMetadata
}

func (a *Application) GetMiddlewares() []nethttp.StrictHTTPMiddlewareFunc {
	return []nethttp.StrictHTTPMiddlewareFunc{a.Middlewares.Auth.TokenToContext, a.Middlewares.RequestMetadata.Time}
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) addLogger() {
	logFile, err := os.OpenFile(a.Options.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
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

func (a *Application) addDB() {
	db, err := sql.Open("pgx", a.Options.DBConnectionString)
	if err != nil {
		panic(err)
	}
	a.DB = db
}

func (a *Application) readConfiguration() {
	opt, err := configuration.ParseOptions("appsettings.json")
	if err != nil {
		panic("read of config failed")
	}
	a.Options = opt
}

func (a *Application) addServices() {
	a.Services = &Services{}
	a.Services.JwtService = auth.NewJwtService(a.Options)
	a.Services.HashingService = auth.NewHashingService()
	a.Services.AuthService = auth.NewAuthService(a.Stores.UserStore, a.Services.HashingService, a.Services.JwtService)
}

func (a *Application) addStores() {
	a.Stores = &Stores{}
	a.Stores.FoodItemStore = fooditem.NewStore(a.DB, a.Logger)
	a.Stores.UserStore = user.NewStore(a.DB, a.Logger)
	a.Stores.MealStore = meal.NewStore(a.DB, a.Logger)
	a.Stores.RecipeStore = recipe.NewStore(a.DB, a.Logger)
}

func (a *Application) addMiddlewares() {
	a.Middlewares = &Middlewares{}
	a.Middlewares.RequestMetadata = middleware.NewRequestMetadata(a.Logger)
	a.Middlewares.Auth = middleware.NewAuth(a.Logger, a.Services.JwtService)
}

func (a *Application) addEndpoints() {
	a.Endpoints = &Endpoints{
		FoodItemEndpoint: fooditem.NewEndpoint(a.Stores.FoodItemStore, a.Logger),
		RecipeEndpoint:   recipe.NewEndpoint(a.Stores.RecipeStore, a.Logger),
		AuthEndpoint:     auth.NewEndpoint(a.Services.AuthService, a.Logger),
		MealEndpoint:     meal.NewEndpoint(a.Stores.MealStore, a.Logger),
	}
}

func (a *Application) staticFS() http.Handler {
	fs := http.FileServer(http.Dir(a.Options.StaticFilesDirectory))
	notFoundInterceptor := middleware.NewFileNotFoundInterceptor(a.Logger)
	return notFoundInterceptor.RespondWithFallback(fs, "/")
}

func (a *Application) rootMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", a.staticFS())

	return mux
}

func (a *Application) Setup() {
	a.readConfiguration()
	a.addLogger()
	a.addDB()
	a.addStores()
	a.addServices()
	a.addMiddlewares()
	a.addEndpoints()

	server := NewServer(a.Endpoints.RecipeEndpoint, a.Endpoints.MealEndpoint, a.Endpoints.FoodItemEndpoint, a.Endpoints.AuthEndpoint)
	s := api.NewStrictHandler(server, a.GetMiddlewares())
	mux := a.rootMux()
	h := api.HandlerFromMux(s, mux)
	a.Server = http.Server{
		Addr:                         a.Options.ListenAddress,
		Handler:                      h,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
}

func (a *Application) Run() {
	a.Logger.Info("Listening", slog.String("address", "http://localhost"), slog.String("port", a.Options.ListenAddress))
	err := a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Error("failure to listen and serve", slog.Any("err", err))
	}
}
