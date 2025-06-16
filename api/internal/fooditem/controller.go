package fooditem

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
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
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
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
	newItem := c.store.Add(item)
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(newItem.ToFoodItemResponse())
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	items := c.store.Get(userID)
	responses := make([]FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToFoodItemResponse())
	}
	enc := json.NewEncoder(w)
	enc.Encode(responses)
}
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {

	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		c.logger.Error("Failed to parse id from path", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := c.store.GetByID(id, userID)
	if err != nil {
		c.logger.Info("fooditem not found", slog.Any("err", err))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc := json.NewEncoder(w)
	enc.Encode(item.ToFoodItemResponse())
}
