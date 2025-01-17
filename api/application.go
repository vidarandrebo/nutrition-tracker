package api

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/user"
)

type Application struct {
	DB            *sql.DB
	FoodItemStore *fooditem.Store
	UserStore     *user.Store
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
