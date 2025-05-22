package internal

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/meal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/middleware"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	Server      http.Server
	DB          *sql.DB
	Options     *configuration.Options
	Logger      *slog.Logger
	Services    *Services
	Stores      *Stores
	Controllers *Controllers
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
}

type Controllers struct {
	FoodItemController *fooditem.Controller
	AuthController     *auth.Controller
	MealController     *meal.Controller
}

type Middlewares struct {
	Auth            *middleware.Auth
	RequestMetadata *middleware.RequestMetadata
	HeaderWriter    *middleware.HeaderWriter
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
	a.Services.JwtService = auth.NewJwtService()
	a.Services.HashingService = auth.NewHashingService()
	a.Services.AuthService = auth.NewAuthService(a.Stores.UserStore, a.Services.HashingService)
}
func (a *Application) addStores() {
	a.Stores = &Stores{}
	a.Stores.FoodItemStore = fooditem.NewStore(a.DB)
	a.Stores.UserStore = user.NewStore(a.DB, a.Logger)
	a.Stores.MealStore = meal.NewStore(a.DB, a.Logger)
}

func (a *Application) addControllers() {
	a.Controllers = &Controllers{}
	a.Controllers.FoodItemController = fooditem.NewController(a.Stores.FoodItemStore, a.Logger)
	a.Controllers.AuthController = auth.NewController(a.Services.AuthService, a.Logger)
	a.Controllers.MealController = meal.NewController(a.Stores.MealStore, a.Logger)
}

func (a *Application) addMiddlewares() {
	a.Middlewares = &Middlewares{}
	a.Middlewares.RequestMetadata = middleware.NewRequestMetadata(a.Logger)
	a.Middlewares.Auth = middleware.NewAuth(a.Logger, a.Services.JwtService)
	a.Middlewares.HeaderWriter = middleware.NewHeaderWriter(a.Logger)
}

func (a *Application) foodItemRoutes() http.Handler {
	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(a.Middlewares.Auth.TokenToContext)
	mw := mwBuilder.Build()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/food-items", a.Controllers.FoodItemController.List)
	mux.HandleFunc("GET /api/food-items/{id}", a.Controllers.FoodItemController.Get)
	mux.HandleFunc("POST /api/food-items", a.Controllers.FoodItemController.Post)
	return mw(mux)
}

func (a *Application) mealRoutes() http.Handler {
	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(a.Middlewares.Auth.TokenToContext)
	mw := mwBuilder.Build()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/meals", a.Controllers.MealController.Post)
	mux.HandleFunc("GET /api/meals", a.Controllers.MealController.Get)
	mux.HandleFunc("GET /api/meals/{id}", a.Controllers.MealController.GetByID)
	mux.HandleFunc("POST /api/meals/{id}/entries", a.Controllers.MealController.PostEntry)
	return mw(mux)
}

func (a *Application) apiMux() http.Handler {
	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(a.Middlewares.RequestMetadata.Time)
	mwBuilder.AddMiddleware(a.Middlewares.HeaderWriter.WriteHeaders)
	mw := mwBuilder.Build()

	mux := http.NewServeMux()
	foodItemRoutes := a.foodItemRoutes()
	mealRoutes := a.mealRoutes()
	mux.Handle("/api/food-items/", foodItemRoutes)
	mux.Handle("/api/food-items", foodItemRoutes)
	mux.Handle("/api/meals/", mealRoutes)
	mux.Handle("/api/meals", mealRoutes)
	mux.HandleFunc("POST /api/login", a.Controllers.AuthController.Login)
	mux.HandleFunc("POST /api/register", a.Controllers.AuthController.Register)
	return mw(mux)
}

func (a *Application) staticFS() http.Handler {
	fs := http.FileServer(http.Dir(a.Options.StaticFilesDirectory))
	notFoundInterceptor := middleware.NewFileNotFoundInterceptor(a.Logger)
	return notFoundInterceptor.RespondWithFallback(fs, "/")
}

func (a *Application) rootMux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", a.staticFS())
	mux.Handle("/api/", a.apiMux())

	return mux
}
func (a *Application) Setup() {
	a.readConfiguration()
	a.addLogger()
	a.addDB()
	a.addStores()
	a.addServices()
	a.addMiddlewares()
	a.addControllers()

	a.Server = http.Server{
		Addr:                         a.Options.ListenAddress,
		Handler:                      a.rootMux(),
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
