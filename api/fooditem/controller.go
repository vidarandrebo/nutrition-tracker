package fooditem

import (
	"fmt"
	"net/http"
)

type FoodItemController struct {
	store *Store
}

func NewFoodItemController(store *Store) *FoodItemController {
	return &FoodItemController{store: store}
}

func (fc *FoodItemController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there")
}
