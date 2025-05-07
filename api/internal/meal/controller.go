package meal

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
)

type Controller struct {
	store *Store
	last  int64
}

func NewController(store *Store) *Controller {
	return &Controller{store: store, last: 0}
}
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	_, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	request, err := utils.ParseJson[PostMealRequest](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	meal := c.store.Add(Meal{
		SequenceNumber: c.last,
		Timestamp:      request.TimeStamp,
	})
	c.last++

	w.WriteHeader(http.StatusCreated)

	response := MealResponse{
		ID:      meal.ID,
		Entries: nil,
	}

	enc := json.NewEncoder(w)
	enc.Encode(response)
}

//func (fc *Controller) List(w http.ResponseWriter, r *http.Request) {
//
//	items := fc.store.Get()
//	responses := make([]FoodItemResponse, 0)
//
//	for _, item := range items {
//		responses = append(responses, item.ToFoodItemResponse())
//	}
//	enc := json.NewEncoder(w)
//	enc.Encode(responses)
//}
