package fooditem

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
}

func NewEndpoint(store *Store, logger *slog.Logger) *Endpoint {
	e := Endpoint{store: store}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) GetApiFoodItems(ctx context.Context, request api.GetApiFoodItemsRequestObject) (api.GetApiFoodItemsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	items := e.store.Get(userID)
	responses := make([]api.FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToFoodItemResponse())
	}
	return api.GetApiFoodItems200JSONResponse(responses), nil
}

func (e Endpoint) PostApiFoodItems(ctx context.Context, request api.PostApiFoodItemsRequestObject) (api.PostApiFoodItemsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	r := PostFoodItemRequest{
		Manufacturer: request.Body.Manufacturer,
		Product:      request.Body.Product,
		Protein:      request.Body.Protein,
		Carbohydrate: request.Body.Carbohydrate,
		Fat:          request.Body.Fat,
		Public:       false,
	}
	if request.Body.KCal == nil {
		r.KCal = 0.0
	} else {
		r.KCal = *request.Body.KCal
	}
	item := r.ToFoodItem()
	item.OwnerID = userID
	newItem := e.store.Add(item)
	return api.PostApiFoodItems201JSONResponse(newItem.ToFoodItemResponse()), nil
}

func (e Endpoint) GetApiFoodItemsId(ctx context.Context, request api.GetApiFoodItemsIdRequestObject) (api.GetApiFoodItemsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	item, err := e.store.GetByID(request.Id)
	if err != nil || !item.HasAccess(userID) {
		e.logger.Info("fooditem not found", slog.Any("err", err))
		return nil, err
	}

	return api.GetApiFoodItemsId200JSONResponse(item.ToFoodItemResponse()), nil
}

func (e Endpoint) DeleteApiFoodItemsId(ctx context.Context, request api.DeleteApiFoodItemsIdRequestObject) (api.DeleteApiFoodItemsIdResponseObject, error) {
	// TODO implement me
	panic("implement me")
}
