package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/middleware"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fileName := filepath.Join("./", "server.log")
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

	app := NewApplication()
	defer app.CloseDB()

	app.FoodItemStore = fooditem.NewStore(app.DB)
	userStore := user.NewStore(app.DB, logger)
	hashingService := auth.NewHashingService()
	app.AuthService = auth.NewAuthService(userStore, hashingService)
	requestTimerMW := middleware.NewRequestTimer(logger)

	mwBuilder := middleware.NewMiddlewareBuilder()
	mwBuilder.AddMiddleware(requestTimerMW.Time)
	mw := mwBuilder.Build()

	mux := http.NewServeMux()

	// Create controller instances
	fs := http.FileServer(http.Dir("./static"))
	foodItemController := fooditem.NewController(app.FoodItemStore)
	userController := auth.NewController(app.AuthService, logger)

	mux.Handle("/", fs)
	mux.HandleFunc("GET /api/fooditems", foodItemController.ListFoodItems)
	mux.HandleFunc("POST /api/login", userController.Login)
	mux.HandleFunc("POST /api/register", userController.Register)

	server := http.Server{
		Addr:                         ":8080",
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

	logger.Info("Listening on http://localhost:8080")
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("failure to listen and serve", slog.Any("err", err))
	}
}
