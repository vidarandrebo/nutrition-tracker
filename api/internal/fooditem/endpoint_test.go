package fooditem_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
)

func TestDeleteApiFoodItemsIdUsesContextUser(t *testing.T) {
	service := NewMockIService(t)
	service.EXPECT().Delete(int64(7), int64(42)).Return(nil)
	endpoint := fooditem.NewEndpoint(service, slog.Default())
	ctx := context.WithValue(context.Background(), "user", &user.User{ID: 42})

	response, err := endpoint.DeleteApiFoodItemsId(ctx, api.DeleteApiFoodItemsIdRequestObject{Id: 7})

	assert.NoError(t, err)
	assert.IsType(t, api.DeleteApiFoodItemsId204Response{}, response)
}

func TestDeleteApiFoodItemsIdWithoutContextUser(t *testing.T) {
	service := NewMockIService(t)
	endpoint := fooditem.NewEndpoint(service, slog.Default())

	response, err := endpoint.DeleteApiFoodItemsId(context.Background(), api.DeleteApiFoodItemsIdRequestObject{Id: 7})

	assert.NoError(t, err)
	assert.IsType(t, api.DeleteApiFoodItemsId401JSONResponse{}, response)
}
