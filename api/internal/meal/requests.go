package meal

import (
	"errors"
	"time"
)

type PostMealRequest struct {
	Timestamp time.Time `json:"timestamp"`
}

type PostMealEntryRequest struct {
	FoodItemID int64   `json:"foodItemId"`
	RecipeID   int64   `json:"recipeId"`
	Amount     float64 `json:"amount"`
}

func (mer PostMealEntryRequest) Validate() (bool, error) {
	if !mer.eitherFoodItemOrRecipe() {
		return false, errors.New("a recipe or foodItem ID needs to be provided")
	}
	if mer.Amount <= 0.0 {
		return false, errors.New("amount needs to be more than 0")
	}
	return true, nil
}

func (mer PostMealEntryRequest) eitherFoodItemOrRecipe() bool {
	return (mer.RecipeID == 0 && mer.FoodItemID != 0) || (mer.RecipeID != 0 && mer.FoodItemID == 0)
}
