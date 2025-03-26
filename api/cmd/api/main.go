package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
)

func main() {

	app := NewApplication()

	app.Setup()

	app.Run()
}
