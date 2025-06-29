package recipe

import (
	"context"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Endpoint struct {
}

func (e Endpoint) GetApiRecipes(ctx context.Context, request api.GetApiRecipesRequestObject) (api.GetApiRecipesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) PostApiRecipes(ctx context.Context, request api.PostApiRecipesRequestObject) (api.PostApiRecipesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
