package meal

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
	"net/http"
	"time"
)

type Controller struct {
	store  *Store
	logger *slog.Logger
	last   int64
}

func NewController(store *Store, logger *slog.Logger) *Controller {
	return &Controller{store: store, last: 0, logger: logger}
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

	response := MealResponse{
		ID:             meal.ID,
		SequenceNumber: meal.SequenceNumber,
		Timestamp:      meal.Timestamp,
		Entries:        nil,
	}

	enc := json.NewEncoder(w)
	enc.Encode(response)
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
		entries := make([]EntryResponse, 0, len(m.Entries))
		for _, e := range m.Entries {
			entries = append(entries, EntryResponse{
				ID:         e.ID,
				Amount:     e.Amount,
				FoodItemID: e.FoodItemID,
			})
		}
		response = append(response, MealResponse{
			ID:             m.ID,
			SequenceNumber: m.SequenceNumber,
			Timestamp:      m.Timestamp,
			Entries:        entries,
		})
	}
	enc := json.NewEncoder(w)
	enc.Encode(response)
}
