package fooditem_test

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

func TestService_Get(t *testing.T) {
	repo := NewMockIRepository(t)
	repo.EXPECT().Get(mock.Anything).Return(make([]fooditem.TableFoodItem, 0), nil)
	repo.EXPECT().GetMicronutrients(mock.Anything).Return(make([]fooditem.TableFoodItemMacronutrient, 0), nil).Maybe()
	repo.EXPECT().GetPortionSizes(mock.Anything).Return(make([]fooditem.TableFoodItemPortionSize, 0), nil).Maybe()
	service := fooditem.NewService(repo, slog.Default())
	_, err := service.Get(1)
	assert.Nil(t, err)
}

func TestService_CheckOwnership(t *testing.T) {
	queryErr := errors.New("database unavailable")
	tests := []struct {
		name    string
		item    fooditem.TableFoodItem
		getErr  error
		ownerID int64
		wantErr error
	}{
		{name: "owned", item: fooditem.TableFoodItem{ID: 1, OwnerID: 2}, ownerID: 2},
		{name: "not owned", item: fooditem.TableFoodItem{ID: 1, OwnerID: 2}, ownerID: 3, wantErr: utils.ErrForbidden},
		{name: "missing", getErr: sql.ErrNoRows, ownerID: 2, wantErr: utils.ErrNotFound},
		{name: "query error", getErr: queryErr, ownerID: 2, wantErr: utils.ErrUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMockIRepository(t)
			repo.EXPECT().GetByID(int64(1)).Return(tt.item, tt.getErr)
			var logs bytes.Buffer
			service := fooditem.NewService(repo, slog.New(slog.NewTextHandler(&logs, nil)))

			err := service.CheckOwnership(1, tt.ownerID)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
			if tt.name == "query error" {
				assert.True(t, strings.Contains(logs.String(), queryErr.Error()), "missing original error in log: %s", logs.String())
				assert.Contains(t, logs.String(), "foodItemID=1")
			}
		})
	}
}

func TestService_InsertNoRowsIsUnknown(t *testing.T) {
	for _, name := range []string{"food item", "nested micronutrient", "nested portion size", "micronutrient", "portion size"} {
		t.Run(name, func(t *testing.T) {
			repo := NewMockIRepository(t)
			insertErr := fmt.Errorf("insert returned nothing: %w", sql.ErrNoRows)
			owned := fooditem.TableFoodItem{ID: 1, OwnerID: 2}
			item := &fooditem.FoodItem{OwnerID: 2}
			switch name {
			case "food item":
				repo.EXPECT().Add(mock.Anything).Return(fooditem.TableFoodItem{}, insertErr)
			case "nested micronutrient":
				item.Micronutrients = []*fooditem.Micronutrient{{}}
				repo.EXPECT().Add(mock.Anything).Return(owned, nil)
				repo.EXPECT().AddMicronutrient(mock.Anything).Return(fooditem.TableFoodItemMacronutrient{}, insertErr)
			case "nested portion size":
				item.PortionSizes = []*fooditem.PortionSize{{}}
				repo.EXPECT().Add(mock.Anything).Return(owned, nil)
				repo.EXPECT().AddPortionSize(mock.Anything).Return(fooditem.TableFoodItemPortionSize{}, insertErr)
			case "micronutrient":
				repo.EXPECT().GetByID(int64(1)).Return(owned, nil)
				repo.EXPECT().AddMicronutrient(mock.Anything).Return(fooditem.TableFoodItemMacronutrient{}, insertErr)
			case "portion size":
				repo.EXPECT().GetByID(int64(1)).Return(owned, nil)
				repo.EXPECT().AddPortionSize(mock.Anything).Return(fooditem.TableFoodItemPortionSize{}, insertErr)
			}

			var logs bytes.Buffer
			service := fooditem.NewService(repo, slog.New(slog.NewTextHandler(&logs, nil)))
			var err error
			switch name {
			case "food item", "nested micronutrient", "nested portion size":
				_, err = service.Add(item)
			case "micronutrient":
				_, err = service.AddMicronutrient(&fooditem.Micronutrient{}, 1, 2)
			case "portion size":
				_, err = service.AddPortionSize(&fooditem.PortionSize{}, 1, 2)
			}
			assert.ErrorIs(t, err, utils.ErrUnknown)
			assert.NotErrorIs(t, err, utils.ErrNotFound)
			assert.Contains(t, logs.String(), sql.ErrNoRows.Error())
		})
	}
}

func TestService_GetByIDNoRowsIsNotFound(t *testing.T) {
	repo := NewMockIRepository(t)
	repo.EXPECT().GetByID(int64(1)).Return(fooditem.TableFoodItem{}, fmt.Errorf("lookup: %w", sql.ErrNoRows))
	service := fooditem.NewService(repo, slog.Default())
	_, err := service.GetByID(1)
	assert.ErrorIs(t, err, utils.ErrNotFound)
}

func TestService_ListNoRowsIsUnknown(t *testing.T) {
	repo := NewMockIRepository(t)
	repo.EXPECT().Get(int64(2)).Return(nil, sql.ErrNoRows)
	service := fooditem.NewService(repo, slog.Default())
	_, err := service.Get(2)
	assert.ErrorIs(t, err, utils.ErrUnknown)
}
