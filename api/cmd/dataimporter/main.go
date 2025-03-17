package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/matvaretabellen"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {

		fmt.Errorf("only one file allowed")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	fileName := filepath.Join("./", "dataimporter.log")
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
	envFile, err := os.Open("./.env")
	env := utils.ReadEnv(envFile)
	envFile.Close()

	app := NewApplication(env)
	app.FoodItemStore = fooditem.NewStore(app.DB)
	userStore := user.NewStore(app.DB, logger)
	hashingService := auth.NewHashingService()
	app.AuthService = auth.NewAuthService(userStore, hashingService)
	defer app.CloseDB()

	matvareTabellenUser, err := userStore.GetUserByEmail("post@matvaretabellen.no")
	if err != nil {
		app.AuthService.RegisterUser(&auth.RegisterRequest{
			Email:    "post@matvaretabellen.no",
			Password: env["MATVARETABELLEN_PASSWORD"],
		})
	}
	matvareTabellenUser, err = userStore.GetUserByEmail("post@matvaretabellen.no")

	foods, err := utils.ParseJson[matvaretabellen.Foods](file)

	if err == nil {
		for _, item := range foods.Items {
			foodItem := fooditem.FromMatvareTabellen(item)
			foodItem.OwnerID = matvareTabellenUser.ID
			fmt.Println(foodItem.Product, "Protein:", foodItem.Protein, "Carbo:", foodItem.Carbohydrate, "Fat:", foodItem.Fat)
			app.FoodItemStore.AddFoodItem(foodItem)
		}
	}
}
