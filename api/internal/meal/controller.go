package meal

import (
	"fmt"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
	"time"
)

type Controller struct {
	store *Store
}

func NewController(store *Store) *Controller {
	return &Controller{store: store}
}
func (fc *Controller) PostMeal(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusCreated)

	fmt.Println(request.TimeStamp.Add(5 * time.Hour))

	//enc := json.NewEncoder(w)
}

//func (fc *Controller) ListFoodItems(w http.ResponseWriter, r *http.Request) {
//
//	items := fc.store.GetFoodItems()
//	responses := make([]FoodItemResponse, 0)
//
//	for _, item := range items {
//		responses = append(responses, item.ToFoodItemResponse())
//	}
//	enc := json.NewEncoder(w)
//	enc.Encode(responses)
//}
