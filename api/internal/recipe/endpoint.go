package recipe

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type Endpoint struct {
	logger  *slog.Logger
	service IService
}

func NewEndpoint(service IService, logger *slog.Logger) *Endpoint {
	e := Endpoint{
		service: service,
	}
	e.logger = logger.With(slog.Any("module", reflect.TypeOf(e)))
	return &e
}

func (e Endpoint) GetApiRecipes(ctx context.Context, request api.GetApiRecipesRequestObject) (api.GetApiRecipesResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, utils.ErrUnauthorized
	}

	recipes, err := e.service.Get(userID)
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

	item := NewRecipe().FromPost(request.Body)
	item.OwnerID = userID

	newItem, err := e.service.Add(item)
	if err != nil {
		return nil, err
	}

	return api.PostApiRecipes201JSONResponse(newItem.ToResponse()), nil
}

func (e Endpoint) DeleteApiRecipesId(ctx context.Context, request api.DeleteApiRecipesIdRequestObject) (api.DeleteApiRecipesIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = e.service.Delete(request.Id, userID)
	if err != nil {
		return api.DeleteApiRecipesId409Response{}, nil
	}

	return api.DeleteApiRecipesId204Response{}, nil
}
