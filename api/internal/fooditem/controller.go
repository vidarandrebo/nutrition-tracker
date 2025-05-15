package fooditem

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
	"net/http"
	"reflect"
)

type Controller struct {
	store  *Store
	logger *slog.Logger
}

func NewController(store *Store, logger *slog.Logger) *Controller {
	c := &Controller{store: store}
	c.logger = logger.With("module", reflect.TypeOf(c))
	return c
}
func (fc *Controller) Post(w http.ResponseWriter, r *http.Request) {
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
	item := request.ToFoodItem()
	item.OwnerID = userID
	newItem := fc.store.Add(item)
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(newItem.ToFoodItemResponse())
}

func (fc *Controller) List(w http.ResponseWriter, r *http.Request) {

	items := fc.store.Get()
	responses := make([]FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToFoodItemResponse())
	}
	enc := json.NewEncoder(w)
	enc.Encode(responses)
}
