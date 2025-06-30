package meal

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Endpoint struct {
	store  *Store
	logger *slog.Logger
	last   int64
}

func NewEndpoint(store *Store, logger *slog.Logger) *Endpoint {
	e := Endpoint{store: store, last: 0}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) GetApiMeals(ctx context.Context, request api.GetApiMealsRequestObject) (api.GetApiMealsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	//	queryValues := r.URL.Query()
	//	dateFrom := queryValues.Get("dateFrom")
	//	dateTo := queryValues.Get("dateTo")
	//
	//	timeTo, err := time.Parse(time.RFC3339, dateTo)
	//	if err != nil {
	//		errs = append(errs, err)
	//	}
	//	timeFrom, err := time.Parse(time.RFC3339, dateFrom)
	//	if err != nil {
	//		errs = append(errs, err)
	//	}
	//
	//	if len(errs) > 0 {
	//		w.WriteHeader(http.StatusBadRequest)
	//		return
	//	}

	meals := e.store.GetByDate(userID, *request.Params.DateFrom, *request.Params.DateTo)

	//	e.logger.Info("meal times", slog.String("from", dateFrom), slog.String("to", dateTo))
	//	e.logger.Info("meal times", slog.Time("from", timeFrom), slog.Time("to", timeTo))

	response := make([]api.MealResponse, 0, len(meals))

	for _, m := range meals {
		response = append(response, m.ToResponse())
	}
	return api.GetApiMeals200JSONResponse(response), nil
}

func (e Endpoint) PostApiMeals(ctx context.Context, request api.PostApiMealsRequestObject) (api.PostApiMealsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	meal := e.store.Add(Meal{
		SequenceNumber: e.last,
		Timestamp:      *request.Body.Timestamp,
		OwnerID:        userID,
	})
	e.last++

	return api.PostApiMeals201JSONResponse(meal.ToResponse()), nil
}

func (e Endpoint) GetApiMealsId(ctx context.Context, request api.GetApiMealsIdRequestObject) (api.GetApiMealsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	meal, err := e.store.GetById(request.Id, userID)
	if err != nil {
		return nil, err
	}

	return api.GetApiMealsId200JSONResponse(meal.ToResponse()), nil
}

func (e Endpoint) PostApiMealsIdEntries(ctx context.Context, request api.PostApiMealsIdEntriesRequestObject) (api.PostApiMealsIdEntriesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	r := PostMealEntryRequest{
		FoodItemID: *request.Body.FoodItemId,
		RecipeID:   *request.Body.RecipeId,
		Amount:     *request.Body.Amount,
	}
	if ok, err := r.Validate(); !ok {
		return nil, err
	}

	entry, err := e.store.AddMealEntry(
		EntryFromRequest(r),
		request.Id,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return api.PostApiMealsIdEntries201JSONResponse(entry.ToResponse()), nil
}
