package fooditem

import (
	"context"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Endpoint struct {
}

func (e Endpoint) GetApiFoodItems(ctx context.Context, request api.GetApiFoodItemsRequestObject) (api.GetApiFoodItemsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) PostApiFoodItems(ctx context.Context, request api.PostApiFoodItemsRequestObject) (api.PostApiFoodItemsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
