package meal

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Controller struct {
	store  *Store
	logger *slog.Logger
	last   int64
}

func NewController(store *Store, logger *slog.Logger) *Controller {
	c := &Controller{store: store, last: 0}
	c.logger = logger.With("module", reflect.TypeOf(c))
	return c
}
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.UserIDFromCtx(r.Context())
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
		Timestamp:      request.Timestamp,
		OwnerID:        userID,
	})
	c.last++

	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	enc.Encode(meal.ToResponse())
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.UserIDFromCtx(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	errs := make([]error, 0)

	queryValues := r.URL.Query()
	dateFrom := queryValues.Get("dateFrom")
	dateTo := queryValues.Get("dateTo")

	timeTo, err := time.Parse(time.RFC3339, dateTo)
	if err != nil {
		errs = append(errs, err)
	}
	timeFrom, err := time.Parse(time.RFC3339, dateFrom)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	meals := c.store.GetByDate(userID, timeFrom, timeTo)

	c.logger.Info("meal times", slog.String("from", dateFrom), slog.String("to", dateTo))
	c.logger.Info("meal times", slog.Time("from", timeFrom), slog.Time("to", timeTo))

	response := make([]MealResponse, 0, len(meals))

	for _, m := range meals {
		response = append(response, m.ToResponse())
	}
	enc := json.NewEncoder(w)
	enc.Encode(response)
}
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
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
	c.logger.Info("get by id", slog.Int64("id", id))

	meal, err := c.store.GetById(id, userID)

	if err != nil {
		enc.Encode(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc.Encode(meal.ToResponse())
}
func (c *Controller) PostEntry(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	request, err := utils.ParseJson[PostMealEntryRequest](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mealID, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	entry, err := c.store.AddMealEntry(Entry{
		FoodItemID: request.FoodItemID,
		Amount:     request.Amount,
	}, mealID, userID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	enc.Encode(entry.ToResponse())
}
