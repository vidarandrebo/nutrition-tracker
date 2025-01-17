package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/vidarandrebo/nutrition-tracker/api"
	"github.com/vidarandrebo/nutrition-tracker/api/fooditem"
	"log"
	"net/http"
)

func main() {
	app := api.NewApplication()
	defer app.CloseDB()

	store := fooditem.NewStore(app.DB)

	foodItem := store.GetFoodItem()

	fmt.Println(foodItem)

	//	fs := http.FileServer(http.Dir("./static"))
	//
	//	mux := http.NewServeMux()
	//	mux.Handle("/", fs)
	//
	//	mux.Handle("/home", &homeHandler{})
	//
	//	log.Print("Listening on localhost:8080")
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
