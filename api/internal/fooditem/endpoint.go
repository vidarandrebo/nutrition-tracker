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
	service IService
	logger  *slog.Logger
}

func NewEndpoint(service IService, logger *slog.Logger) *Endpoint {
	e := Endpoint{service: service}
	e.logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) GetApiFoodItems(ctx context.Context, request api.GetApiFoodItemsRequestObject) (api.GetApiFoodItemsResponseObject, error) {
	user, err := auth.UserFromCtx(ctx)
	if err != nil {
		return api.GetApiFoodItems401JSONResponse{}, nil
	}
	items, err := e.service.Get(user.ID)
	responses := make([]api.FoodItemResponse, 0)

	for _, item := range items {
		responses = append(responses, item.ToResponse())
	}
	return api.GetApiFoodItems200JSONResponse(responses), nil
}

func (e Endpoint) PostApiFoodItems(ctx context.Context, request api.PostApiFoodItemsRequestObject) (api.PostApiFoodItemsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return api.PostApiFoodItems401JSONResponse{}, nil
	}

	e.logger.Info("new foodItem", slog.Bool("isPublic", request.Body.IsPublic))
	item := NewFoodItem().FromRequest(request.Body)
	item.OwnerID = userID
	newItem, err := e.service.Add(item)
	if err != nil {
		return nil, utils.ErrUnknown
	}

	return api.PostApiFoodItems201JSONResponse(newItem.ToResponse()), nil
}

func (e Endpoint) GetApiFoodItemsId(ctx context.Context, request api.GetApiFoodItemsIdRequestObject) (api.GetApiFoodItemsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return api.GetApiFoodItemsId401JSONResponse{}, nil
	}
	item, err := e.service.GetByID(request.Id)
	if err != nil {
		return api.GetApiFoodItemsId404JSONResponse{}, nil
	}
	if !item.HasAccess(userID) {
		return api.GetApiFoodItemsId403JSONResponse{}, nil
	}

	return api.GetApiFoodItemsId200JSONResponse(item.ToResponse()), nil
}

func (e Endpoint) DeleteApiFoodItemsId(ctx context.Context, request api.DeleteApiFoodItemsIdRequestObject) (api.DeleteApiFoodItemsIdResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return api.DeleteApiFoodItemsId401JSONResponse{}, nil
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
		return api.PostApiFoodItemsIdPortions401JSONResponse{}, nil
	}
	ps, err := e.service.AddPortionSize(FromPortionSizePost(request.Body), request.Id, userID)

	if errors.Is(err, utils.ErrEntityNotFound) {
		return api.PostApiFoodItemsIdPortions404JSONResponse{}, nil
	} else if errors.Is(err, utils.ErrEntityNotOwned) {
		return api.PostApiFoodItemsIdPortions403JSONResponse{}, nil
	} else if err != nil {
		return nil, utils.ErrUnknown
	}
	return api.PostApiFoodItemsIdPortions201JSONResponse(ps.ToResponse()), nil
}

func (e Endpoint) PostApiFoodItemsIdMicronutrients(ctx context.Context, request api.PostApiFoodItemsIdMicronutrientsRequestObject) (api.PostApiFoodItemsIdMicronutrientsResponseObject, error) {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return api.PostApiFoodItemsIdMicronutrients401JSONResponse{}, nil
	}

	micronutrient := FromMicronutrientPost(request.Body)

	ps, err := e.service.AddMicronutrient(micronutrient, request.Id, userID)
	if errors.Is(err, utils.ErrEntityNotFound) {
		return api.PostApiFoodItemsIdMicronutrients404JSONResponse{}, nil
	} else if errors.Is(err, utils.ErrEntityNotOwned) {
		return api.PostApiFoodItemsIdMicronutrients403JSONResponse{}, nil
	} else if err != nil {
		return nil, utils.ErrUnknown
	}
	return api.PostApiFoodItemsIdMicronutrients201JSONResponse(ps.ToResponse()), nil
}
