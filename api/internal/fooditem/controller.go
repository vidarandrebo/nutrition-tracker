package fooditem

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	store *Store
}

func NewController(store *Store) *Controller {
	return &Controller{store: store}
}

func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {
	responses := make([]GetFoodItemResponse, 0)

	for i := 0; i < 10; i++ {
		responses = append(responses, GetFoodItemResponse{
			ID:           int64(i + 1),
			Manufacturer: "bogus",
			Product:      "aaa",
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
