package meal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostMealEntryRequest_Validate_BothIDs(t *testing.T) {
	r := PostMealEntryRequest{
		FoodItemID: 1,
		RecipeID:   2,
		Amount:     1.0,
	}

	ok, err := r.Validate()

	assert.NotNil(t, err)
	assert.False(t, ok)
}

func TestPostMealEntryRequest_Validate_ZeroAmount(t *testing.T) {
	r := PostMealEntryRequest{
		FoodItemID: 0,
		RecipeID:   2,
		Amount:     0.0,
	}

	ok, _ := r.Validate()
	assert.False(t, ok)
}

func TestPostMealEntryRequest_Validate_Success(t *testing.T) {
	r := PostMealEntryRequest{
		FoodItemID: 2,
		RecipeID:   0,
		Amount:     1.0,
	}

	ok, err := r.Validate()

	assert.Nil(t, err)
	assert.True(t, ok)
}
