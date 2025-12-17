package meal

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Endpoint struct {
	store  IService
	logger *slog.Logger
	last   int64
}

func (e Endpoint) PostApiMealsIdFoodItemEntries(ctx context.Context, request api.PostApiMealsIdFoodItemEntriesRequestObject) (api.PostApiMealsIdFoodItemEntriesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	entry, err := e.store.AddFoodItemEntry(FIMEFromRequest(*request.Body), request.Id, userID)
	if err != nil {
		return nil, err
	}

	return api.PostApiMealsIdFoodItemEntries201JSONResponse(entry.ToResponse()), nil
}

func (e Endpoint) PostApiMealsIdMacronutrientEntries(ctx context.Context, request api.PostApiMealsIdMacronutrientEntriesRequestObject) (api.PostApiMealsIdMacronutrientEntriesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	entry, err := e.store.AddMacroEntry(MNEFromRequest(*request.Body), request.Id, userID)
	if err != nil {
		return nil, err
	}

	return api.PostApiMealsIdMacronutrientEntries201JSONResponse(entry.ToResponse()), nil
}

func (e Endpoint) PostApiMealsIdRecipeEntries(ctx context.Context, request api.PostApiMealsIdRecipeEntriesRequestObject) (api.PostApiMealsIdRecipeEntriesResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func NewEndpoint(service IService, logger *slog.Logger) *Endpoint {
	e := Endpoint{store: service, last: 0}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) GetApiMeals(ctx context.Context, request api.GetApiMealsRequestObject) (api.GetApiMealsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	meals, err := e.store.GetByDate(*request.Params.DateFrom, *request.Params.DateTo, userID)
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

	meal, err := e.store.Add(&Meal{
		SequenceNumber: 0,
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

func (e Endpoint) DeleteApiMealsId(ctx context.Context, request api.DeleteApiMealsIdRequestObject) (api.DeleteApiMealsIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = e.store.Delete(request.Id, userId)
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
	if userId != 0 {
		panic("")
	}

	//	err = e.store.DeleteMealEntry(request.EntryId, request.MealId, userId)
	//	if err != nil {
	//		return nil, err
	//	}
	return api.DeleteApiMealsMealIdEntriesEntryId204Response{}, nil
}
