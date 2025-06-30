package recipe

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Endpoint struct {
	logger *slog.Logger
	store  *Store
}

func NewEndpoint(store *Store, logger *slog.Logger) *Endpoint {
	e := Endpoint{
		store: store,
	}
	e.logger = logger.With(slog.Any("module", reflect.TypeOf(e)))
	return &e
}

func (e Endpoint) GetApiRecipes(ctx context.Context, request api.GetApiRecipesRequestObject) (api.GetApiRecipesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	recipes, err := e.store.Get(userID)
	if err != nil {
		return nil, err
	}
	responses := make([]api.RecipeResponse, 0, len(recipes))
	for _, recipe := range recipes {
		responses = append(responses, recipe.ToResponse())
	}

	return api.GetApiRecipes200JSONResponse(responses), nil
}

func (e Endpoint) PostApiRecipes(ctx context.Context, request api.PostApiRecipesRequestObject) (api.PostApiRecipesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	entries := make([]Entry, 0, len(*request.Body.Entries))
	for _, e := range *request.Body.Entries {
		entries = append(entries, Entry{
			Amount:     *e.Amount,
			FoodItemID: *e.FoodItemId,
		})
	}
	recipe, err := e.store.Add(Recipe{Name: *request.Body.Name, Entries: entries, OwnerID: userID})
	if err != nil {
		return nil, err
	}
	return api.PostApiRecipes201JSONResponse(recipe.ToResponse()), nil
}
