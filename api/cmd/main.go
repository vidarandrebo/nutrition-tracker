package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/vidarandrebo/nutrition-tracker/api"
	"github.com/vidarandrebo/nutrition-tracker/api/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/user"
	"log"
	"net/http"
)

func main() {
	app := api.NewApplication()
	defer app.CloseDB()

	app.FoodItemStore = fooditem.NewStore(app.DB)

	app.UserStore = user.NewStore(app.DB)

	for i := 0; i < 100; i++ {
		u := user.User{
			ID:           0,
			Name:         "",
			PasswordHash: nil,
		}
		app.UserStore.AddUser(&u)
	}

	users := app.UserStore.ListUsers()

	for _, u := range users {
		fmt.Println(u)
	}

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	mux.Handle("/home", &homeHandler{})
	log.Print("Listening on localhost:8080")

	//	err = http.ListenAndServe("localhost:8080", mux)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}

type homeHandler struct {
}

func (hh *homeHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	numBytes, err := fmt.Fprintf(rw, "hello from this side of the app")
	if err != nil {
		log.Println("something went wrong")
	}
	log.Println("Wrote ", numBytes, " bytes")
}
