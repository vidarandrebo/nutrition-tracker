package internal

import (
	"database/sql"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/consts"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
)

type Application struct {
	DB            *sql.DB
	FoodItemStore *fooditem.Store
	AuthService   *auth.Service
	Env           map[string]string
}

func (a *Application) CloseDB() {
	a.DB.Close()
}

func NewApplication(env map[string]string) *Application {
	connString := env["DB_CONN_STRING"]
	//db, err := sql.Open("pgx", "postgresql://postgres@localhost:5432/nutritiontracker")
	db, err := sql.Open(consts.DatabaseDriverName, connString)
	if err != nil {
		panic(err)
	}

	return &Application{
		DB: db,
	}
}
