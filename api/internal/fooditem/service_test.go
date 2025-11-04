package fooditem_test

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
)

func TestService_Get(t *testing.T) {
	repo := NewMockIRepository(t)
	repo.EXPECT().Get(mock.Anything).Return(make([]fooditem.TableFoodItem, 0), nil)
	repo.EXPECT().GetMicronutrients(mock.Anything).Return(make([]fooditem.TableMicronutrient, 0), nil).Maybe()
	repo.EXPECT().GetPortionSizes(mock.Anything).Return(make([]fooditem.TablePortionSize, 0), nil).Maybe()
	service := fooditem.NewService(repo, slog.Default())
	_, err := service.Get(1)
	assert.Nil(t, err)
}
