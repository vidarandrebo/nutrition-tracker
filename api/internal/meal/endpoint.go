package meal

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Endpoint struct {
	service IService
	logger  *slog.Logger
	last    int64
}

func (e *Endpoint) PostApiMealsMealIdFoodItemEntries(ctx context.Context, request api.PostApiMealsMealIdFoodItemEntriesRequestObject) (api.PostApiMealsMealIdFoodItemEntriesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	entry, err := e.service.AddFoodItemEntry(FIMEFromRequest(*request.Body), request.MealId, userID)
	if err != nil {
		return nil, err
	}

	return api.PostApiMealsMealIdFoodItemEntries201JSONResponse(entry.ToResponse()), nil
}

func (e *Endpoint) PostApiMealsMealIdMacronutrientEntries(ctx context.Context, request api.PostApiMealsMealIdMacronutrientEntriesRequestObject) (api.PostApiMealsMealIdMacronutrientEntriesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	entry, err := e.service.AddMacroEntry(MNEFromRequest(*request.Body), request.MealId, userID)
	if err != nil {
		return nil, err
	}

	return api.PostApiMealsMealIdMacronutrientEntries201JSONResponse(entry.ToResponse()), nil
}

func (e *Endpoint) PostApiMealsMealIdRecipeEntries(ctx context.Context, request api.PostApiMealsMealIdRecipeEntriesRequestObject) (api.PostApiMealsMealIdRecipeEntriesResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

func NewEndpoint(service IService, logger *slog.Logger) *Endpoint {
	e := Endpoint{service: service, last: 0}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e *Endpoint) GetApiMeals(ctx context.Context, request api.GetApiMealsRequestObject) (api.GetApiMealsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	meals, err := e.service.GetByDate(*request.Params.DateFrom, *request.Params.DateTo, userID)
	if err != nil {
		return nil, err
	}

	response := make([]api.MealResponse, 0, len(meals))

	for _, m := range meals {
		response = append(response, m.ToResponse())
	}
	return api.GetApiMeals200JSONResponse(response), nil
}

func (e *Endpoint) PostApiMeals(ctx context.Context, request api.PostApiMealsRequestObject) (api.PostApiMealsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	meal, err := e.service.Add(&Meal{
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

func (e *Endpoint) GetApiMealsId(ctx context.Context, request api.GetApiMealsIdRequestObject) (api.GetApiMealsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	meal, err := e.service.GetById(request.Id, userID)
	if err != nil {
		return nil, err
	}

	return api.GetApiMealsId200JSONResponse(meal.ToResponse()), nil
}

func (e *Endpoint) DeleteApiMealsId(ctx context.Context, request api.DeleteApiMealsIdRequestObject) (api.DeleteApiMealsIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = e.service.Delete(request.Id, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsId204Response{}, nil
}

func (e *Endpoint) DeleteApiMealsMealIdFoodItemEntriesFoodItemEntryId(ctx context.Context, request api.DeleteApiMealsMealIdFoodItemEntriesFoodItemEntryIdRequestObject) (api.DeleteApiMealsMealIdFoodItemEntriesFoodItemEntryIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = e.service.DeleteFoodItemEntry(request.FoodItemEntryId, request.MealId, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsMealIdFoodItemEntriesFoodItemEntryId204Response{}, nil
}

func (e *Endpoint) DeleteApiMealsMealIdMacronutrientEntriesMacronutrientEntryId(ctx context.Context, request api.DeleteApiMealsMealIdMacronutrientEntriesMacronutrientEntryIdRequestObject) (api.DeleteApiMealsMealIdMacronutrientEntriesMacronutrientEntryIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = e.service.DeleteMacronutrientEntry(request.MacronutrientEntryId, request.MealId, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsMealIdMacronutrientEntriesMacronutrientEntryId204Response{}, nil
}

func (e *Endpoint) DeleteApiMealsMealIdRecipeEntriesRecipeEntryId(ctx context.Context, request api.DeleteApiMealsMealIdRecipeEntriesRecipeEntryIdRequestObject) (api.DeleteApiMealsMealIdRecipeEntriesRecipeEntryIdResponseObject, error) {
	userId, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = e.service.DeleteRecipeEntry(request.RecipeEntryId, request.MealId, userId)
	if err != nil {
		return nil, err
	}
	return api.DeleteApiMealsMealIdRecipeEntriesRecipeEntryId204Response{}, nil
}
