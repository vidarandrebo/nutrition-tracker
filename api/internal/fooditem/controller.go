package fooditem

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
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
	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	request, err := utils.ParseJson[PostFoodItemRequest](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newItem := fc.store.AddFoodItem(request, userID)
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(newItem.ToFoodItemResponse())
}

func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {

	items := fc.store.GetFoodItems()
	responses := make([]FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToFoodItemResponse())
	}
	enc := json.NewEncoder(w)
	enc.Encode(responses)
}
