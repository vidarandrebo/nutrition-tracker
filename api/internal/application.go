package internal

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/configuration"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/middleware"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
)

type Application struct {
	Server        http.Server
	DB            *sql.DB
	FoodItemStore *fooditem.Store
	AuthService   *auth.Service
	Env           map[string]string
	Options       *configuration.Options
	Logger        *slog.Logger
}

func (a *Application) CloseDB() {
	a.DB.Close()
}

func NewApplication() *Application {

	return &Application{}
}
func (a *Application) Setup() {
	envFile, err := os.Open("./.env")
	env := utils.ReadEnv(envFile)
	envFile.Close()
	connString := env["DB_CONN_STRING"]
	//db, err := sql.Open("pgx", "postgresql://postgres@localhost:5432/nutritiontracker")
	a.DB, err = sql.Open(configuration.DatabaseDriverName, connString)
	if err != nil {
		panic(err)
	}
	fileName := filepath.Join("/var/log/nutrition-tracker/server.log")
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logHandlerOpts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logWriter := io.MultiWriter(logFile, os.Stderr)
	logHandler := slog.NewTextHandler(logWriter, &logHandlerOpts)

	logger := slog.New(logHandler)

	defer a.CloseDB()

	a.FoodItemStore = fooditem.NewStore(a.DB)
	userStore := user.NewStore(a.DB, logger)
	hashingService := auth.NewHashingService()
	a.AuthService = auth.NewAuthService(userStore, hashingService)
	jwtService := auth.NewJwtService()
	requestTimerMW := middleware.NewRequestTimer(logger)
	authMiddleWare := middleware.NewAuth(logger, jwtService)

	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(requestTimerMW.Time)
	mwBuilder.AddMiddleware(authMiddleWare.TokenToContext)
	mw := mwBuilder.Build()

	mux := http.NewServeMux()

	// Create controller instances
	fs := http.FileServer(http.Dir("./static"))

	notFoundInterceptor := middleware.NewFileNotFoundInterceptor(logger)
	foodItemController := fooditem.NewController(a.FoodItemStore)
	userController := auth.NewController(a.AuthService, logger)

	mux.Handle("/", notFoundInterceptor.RespondWithFallback(fs, "/"))
	mux.HandleFunc("GET /api/food-items", foodItemController.ListFoodItems)
	mux.HandleFunc("POST /api/food-items", foodItemController.PostFoodItem)
	mux.HandleFunc("POST /api/login", userController.Login)
	mux.HandleFunc("POST /api/register", userController.Register)

	a.Server = http.Server{
		Addr:                         ":8081",
		Handler:                      mw(mux),
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
	a.Logger.Info("Listening on http://localhost:8081")
	err := a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Error("failure to listen and serve", slog.Any("err", err))
	}
}
