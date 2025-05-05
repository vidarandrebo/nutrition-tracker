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
	MiddleWares *MiddleWares
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

type MiddleWares struct {
	Auth            *middleware.Auth
	RequestMetadata *middleware.RequestMetadata
}

func (a *Application) CloseDB() {
	a.DB.Close()
}

func NewApplication() *Application {
	return &Application{}
}
func (a *Application) configureLogger() {
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

func (a *Application) configureDB() {
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
func (a *Application) configureServices() {
	a.Services = &Services{}
	a.Services.JwtService = auth.NewJwtService()
	a.Services.HashingService = auth.NewHashingService()
	a.Services.AuthService = auth.NewAuthService(a.Stores.UserStore, a.Services.HashingService)
}
func (a *Application) configureStores() {
	a.Stores = &Stores{}
	a.Stores.FoodItemStore = fooditem.NewStore(a.DB)
	a.Stores.UserStore = user.NewStore(a.DB, a.Logger)
	a.Stores.MealStore = meal.NewStore(a.DB, a.Logger)
}

func (a *Application) configureControllers() {
	a.Controllers = &Controllers{}
	a.Controllers.FoodItemController = fooditem.NewController(a.Stores.FoodItemStore)
	a.Controllers.AuthController = auth.NewController(a.Services.AuthService, a.Logger)
	a.Controllers.MealController = meal.NewController(a.Stores.MealStore)
}

func (a *Application) configureMiddlewares() {
	a.MiddleWares = &MiddleWares{}
	a.MiddleWares.RequestMetadata = middleware.NewRequestMetadata(a.Logger)
	a.MiddleWares.Auth = middleware.NewAuth(a.Logger, a.Services.JwtService)
}

func (a *Application) Setup() {
	a.readConfiguration()
	a.configureLogger()
	a.configureDB()
	a.configureStores()
	a.configureServices()
	a.configureMiddlewares()
	a.configureControllers()

	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(a.MiddleWares.RequestMetadata.Time)
	apiMW := mwBuilder.Build()

	fiMWBuilder := middleware.NewMiddlewareBuilder()
	fiMWBuilder.AddMiddleware(a.MiddleWares.Auth.TokenToContext)
	foodItemMW := fiMWBuilder.Build()

	mealMWBuilder := middleware.NewMiddlewareBuilder()
	mealMWBuilder.AddMiddleware(a.MiddleWares.Auth.TokenToContext)
	mealMW := mealMWBuilder.Build()

	mux := http.NewServeMux()

	// Create controller instances
	fs := http.FileServer(http.Dir(a.Options.StaticFilesDirectory))

	notFoundInterceptor := middleware.NewFileNotFoundInterceptor(a.Logger)

	mux.Handle("/", notFoundInterceptor.RespondWithFallback(fs, "/"))

	foodItemControllerMux := http.NewServeMux()

	foodItemControllerMux.HandleFunc("GET /api/food-items", a.Controllers.FoodItemController.ListFoodItems)
	foodItemControllerMux.HandleFunc("POST /api/food-items", a.Controllers.FoodItemController.PostFoodItem)

	mealControllerMux := http.NewServeMux()
	mealControllerMux.HandleFunc("POST /api/meals", a.Controllers.MealController.PostMeal)

	apiMux := http.NewServeMux()
	apiMux.Handle("/api/food-items", foodItemMW(foodItemControllerMux))
	apiMux.Handle("/api/meals", mealMW(mealControllerMux))
	apiMux.HandleFunc("POST /api/login", a.Controllers.AuthController.Login)
	apiMux.HandleFunc("POST /api/register", a.Controllers.AuthController.Register)

	mux.Handle("/api/", apiMW(apiMux))

	a.Server = http.Server{
		Addr:                         a.Options.ListenAddress,
		Handler:                      mux,
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
