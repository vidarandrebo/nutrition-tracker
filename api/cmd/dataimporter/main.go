package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/vidarandrebo/nutrition-tracker/api/internal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/matvaretabellen"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"golang.org/x/sync/semaphore"
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
	matvareTabellenUser, err := app.Stores.UserRepository.GetByEmail(matvareTabellenCredentials.Email)
	if err != nil {
		app.Services.AuthService.RegisterUser(auth.Register{
			Email:    matvareTabellenCredentials.Email,
			Password: matvareTabellenCredentials.Password,
		})
	}
	matvareTabellenUser, err = app.Stores.UserRepository.GetByEmail("post@matvaretabellen.no")

	foods, err := utils.ParseJson[matvaretabellen.Foods](file)

	hc := http.Client{}

	c, err := api.NewClientWithResponses(app.Options.DataImporterTarget, api.WithHTTPClient(&hc))
	response, err := c.PostApiLoginWithResponse(context.Background(), api.LoginRequest{
		Email:    matvareTabellenCredentials.Email,
		Password: matvareTabellenCredentials.Password,
	})
	bearer := response.JSON200.Token

	reqEdit := func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", bearer))
		return nil
	}
	fmt.Println(response)

	sem := semaphore.NewWeighted(10)
	var wg sync.WaitGroup
	if err == nil {
		for _, item := range foods.Items {
			wg.Add(1)
			sem.Acquire(context.Background(), 1)
			go func() {
				foodItem := fooditem.FromMatvareTabellen(item)
				foodItem.OwnerID = matvareTabellenUser.ID
				fmt.Println(foodItem.Product, "Protein:", foodItem.Protein, "Carbo:", foodItem.Carbohydrate, "Fat:", foodItem.Fat)

				r, err := c.PostApiFoodItemsWithResponse(context.Background(), api.PostFoodItemRequest{
					Carbohydrate: foodItem.Carbohydrate,
					Fat:          foodItem.Fat,
					IsPublic:     foodItem.Public,
					KCal:         &foodItem.KCal,
					Manufacturer: foodItem.Manufacturer,
					Product:      foodItem.Product,
					Protein:      foodItem.Protein,
				}, reqEdit)

				if err == nil {
					for _, mn := range foodItem.Micronutrients {
						c.PostApiFoodItemsIdMicronutrients(context.Background(), r.JSON201.Id, api.PostFoodItemMicronutrient{
							Amount: mn.Amount,
							Name:   mn.Name,
						}, reqEdit)
					}
				}

				wg.Done()
				sem.Release(1)
			}()
		}
	}
	wg.Wait()
}
