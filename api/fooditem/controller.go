package fooditem

import (
	"fmt"
	"net/http"
)

type Controller struct {
	store *Store
}

func NewController(store *Store) *Controller {
	return &Controller{store: store}
}

func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there")
}
