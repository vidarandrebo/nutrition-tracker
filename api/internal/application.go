package internal

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
)

type Application struct {
	DB            *sql.DB
	FoodItemStore *fooditem.Store
	AuthService   *auth.Service
}

func (a *Application) CloseDB() {
	a.DB.Close()
}

func NewApplication() *Application {
	db, err := sql.Open("pgx", "postgresql://postgres@localhost:5432/nutritiontracker")
	if err != nil {
		panic(err)
	}

	return &Application{
		DB: db,
	}
}
