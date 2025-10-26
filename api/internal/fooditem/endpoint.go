package fooditem

import (
	"context"
	"errors"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type Endpoint struct {
	service *IService
	logger  *slog.Logger
}

func NewEndpoint(service *IService, logger *slog.Logger) *Endpoint {
	e := Endpoint{service: service}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) GetApiFoodItems(ctx context.Context, request api.GetApiFoodItemsRequestObject) (api.GetApiFoodItemsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	items := e.service.Get(userID)
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
		Public:       request.Body.IsPublic,
	}
	e.logger.Info("new foodItem", slog.Bool("isPublic", r.Public))
	if request.Body.KCal == nil {
		r.KCal = 0.0
	} else {
		r.KCal = *request.Body.KCal
	}
	item := r.ToFoodItem()
	item.OwnerID = userID
	newItem := e.service.Add(item)
	return api.PostApiFoodItems201JSONResponse(newItem.ToFoodItemResponse()), nil
}

func (e Endpoint) GetApiFoodItemsId(ctx context.Context, request api.GetApiFoodItemsIdRequestObject) (api.GetApiFoodItemsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	item, err := e.service.GetByID(request.Id)
	if err != nil || !item.HasAccess(userID) {
		e.logger.Info("fooditem not found", slog.Any("err", err))
		return nil, err
	}

	return api.GetApiFoodItemsId200JSONResponse(item.ToFoodItemResponse()), nil
}

func (e Endpoint) DeleteApiFoodItemsId(ctx context.Context, request api.DeleteApiFoodItemsIdRequestObject) (api.DeleteApiFoodItemsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	err = e.service.Delete(request.Id, userID)
	if err != nil {
		return api.DeleteApiFoodItemsId409Response{}, nil
	}
	return api.DeleteApiFoodItemsId204Response{}, nil
}

func (e Endpoint) PostApiFoodItemsIdPortions(ctx context.Context, request api.PostApiFoodItemsIdPortionsRequestObject) (api.PostApiFoodItemsIdPortionsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, utils.ErrUnauthorized
	}
	portionSize := PortionSize{
		Amount: request.Body.Amount,
		Name:   request.Body.Name,
	}
	ps, err := e.service.AddPortionSize(request.Id, portionSize, userID)
	if errors.Is(err, utils.ErrEntityNotFound) {
		return api.PostApiFoodItemsIdPortions404Response{}, nil
	} else if errors.Is(err, utils.ErrEntityNotOwned) {
		return nil, utils.ErrEntityNotOwned
	} else if err != nil {
		return nil, utils.ErrUnknown
	}
	return api.PostApiFoodItemsIdPortions201JSONResponse(ps.ToResponse()), nil
}

func (e Endpoint) PostApiFoodItemsIdMicronutrients(ctx context.Context, request api.PostApiFoodItemsIdMicronutrientsRequestObject) (api.PostApiFoodItemsIdMicronutrientsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return nil, utils.ErrUnauthorized
	}
	micronutrient := Micronutrient{
		Amount: request.Body.Amount,
		Name:   request.Body.Name,
	}
	ps, err := e.service.AddMicronutrient(request.Id, micronutrient, userID)
	if errors.Is(err, utils.ErrEntityNotFound) {
		return api.PostApiFoodItemsIdMicronutrients404Response{}, nil
	} else if errors.Is(err, utils.ErrEntityNotOwned) {
		return nil, utils.ErrEntityNotOwned
	} else if err != nil {
		return nil, utils.ErrUnknown
	}
	return api.PostApiFoodItemsIdMicronutrients201JSONResponse(ps.ToResponse()), nil
}
