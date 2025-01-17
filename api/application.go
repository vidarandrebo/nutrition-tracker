package api

import "database/sql"

type Application struct {
	DB *sql.DB
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
