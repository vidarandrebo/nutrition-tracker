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

	meals, err := e.store.GetByDate(userID, *request.Params.DateFrom, *request.Params.DateTo)
	if err != nil {
		return nil, err
	}

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

	meal, err := e.store.Add(Meal{
		SequenceNumber: e.last,
		Timestamp:      request.Body.Timestamp,
		OwnerID:        userID,
	})
	e.last++
	if err != nil {
		return nil, err
	}

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
		Amount:     request.Body.Amount,
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

func (e Endpoint) DeleteApiMealsId(ctx context.Context, request api.DeleteApiMealsIdRequestObject) (api.DeleteApiMealsIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = e.store.DeleteMeal(request.Id, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsId204Response{}, nil
}

func (e Endpoint) DeleteApiMealsMealIdEntriesEntryId(ctx context.Context, request api.DeleteApiMealsMealIdEntriesEntryIdRequestObject) (api.DeleteApiMealsMealIdEntriesEntryIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = e.store.DeleteMealEntry(request.EntryId, request.MealId, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsMealIdEntriesEntryId204Response{}, nil
}
