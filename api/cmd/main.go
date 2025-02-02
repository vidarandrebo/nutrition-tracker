package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"log"
	"net/http"
)

func main() {
	app := NewApplication()
	defer app.CloseDB()

	app.FoodItemStore = fooditem.NewStore(app.DB)
	userStore := user.NewStore(app.DB)
	hashingService := auth.NewHashingService()
	app.AuthService = auth.NewAuthService(userStore, hashingService)

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()

	// Create controller instances
	foodItemController := fooditem.NewController(app.FoodItemStore)
	userController := auth.NewController(app.AuthService)

	mux.Handle("/", fs)
	mux.HandleFunc("GET /api/fooditems", foodItemController.ListFoodItems)
	mux.HandleFunc("POST /api/login", userController.Login)
	mux.HandleFunc("POST /api/register", userController.Register)

	log.Print("Listening on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
