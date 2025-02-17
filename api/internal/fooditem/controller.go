package fooditem

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
)

type Controller struct {
	store *Store
}

func NewController(store *Store) *Controller {
	return &Controller{store: store}
}
func (fc *Controller) PostFoodItem(w http.ResponseWriter, r *http.Request) {
	request, err := utils.ParseJson[PostFoodItemRequest](r.Body)
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := 0

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newItem := fc.store.AddFoodItem(request, int64(id))
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(newItem)
}

func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {
	items := fc.store.GetFoodItems()
	responses := make([]FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToFoodItemResponse())
	}
	enc := json.NewEncoder(w)
	enc.Encode(responses)
	w.WriteHeader(http.StatusOK)
}
