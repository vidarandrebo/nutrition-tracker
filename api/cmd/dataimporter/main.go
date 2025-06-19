package main

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/matvaretabellen"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
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

	app := NewImporter()
	app.Setup()

	matvareTabellenCredentials, ok := app.Options.SystemUsers["Matvaretabellen"]
	if !ok {
		panic("no systemuser config registered for matvaretabellen")
	}
	matvareTabellenUser, err := app.Stores.UserStore.GetUserByEmail(matvareTabellenCredentials.Email)
	if err != nil {
		app.Services.AuthService.RegisterUser(auth.RegisterRequest{
			Email:    matvareTabellenCredentials.Email,
			Password: matvareTabellenCredentials.Password,
		})
	}
	matvareTabellenUser, err = app.Stores.UserStore.GetUserByEmail("post@matvaretabellen.no")

	foods, err := utils.ParseJson[matvaretabellen.Foods](file)

	if err == nil {
		for _, item := range foods.Items {
			foodItem := fooditem.FromMatvareTabellen(item)
			foodItem.OwnerID = matvareTabellenUser.ID
			fmt.Println(foodItem.Product, "Protein:", foodItem.Protein, "Carbo:", foodItem.Carbohydrate, "Fat:", foodItem.Fat)
			app.Stores.FoodItemStore.Add(foodItem)
		}
	}
}
