package fooditem

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
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
	newItem := fc.store.AddFoodItem(request, id)
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(newItem)
}

func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {
	items := fc.store.GetFoodItems()
	responses := make([]GetFoodItemResponse, 0)

	for i, item := range items {

		slog.Info("hello")
		responses = append(responses, GetFoodItemResponse{
			ID:           item.ID,
			Manufacturer: item.Manufacturer,
			Product:      item.Product,
			Macronutrients: GetMacronutrientResponse{
				Protein:      float64(i + 1),
				Carbohydrate: float64(i + 2),
				Fat:          float64(i + 3),
				KCal:         float64(i + 4),
			},
		})
	}
	enc := json.NewEncoder(w)
	enc.Encode(responses)
	w.WriteHeader(http.StatusOK)
}
