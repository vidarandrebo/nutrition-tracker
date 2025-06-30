package meal

import (
	"database/sql"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type Entry struct {
	ID         int64
	foodItemID sql.NullInt64
	recipeID   sql.NullInt64
	Amount     float64
}

func (e Entry) FoodItemID() any {
	if !e.foodItemID.Valid {
		return nil
	}
	return e.foodItemID.Int64
}

func (e Entry) RecipeID() any {
	if !e.recipeID.Valid {
		return nil
	}
	return e.recipeID.Int64
}

func (e Entry) ToResponse() api.MealEntryResponse {
	return api.MealEntryResponse{
		Id:         &e.ID,
		Amount:     &e.Amount,
		FoodItemId: &e.foodItemID.Int64,
		RecipeId:   &e.recipeID.Int64,
	}
}

func (e Entry) IsValid() bool {
	return (e.ID != 0) && (e.Amount != 0.0)
}

func EntryFromRequest(r PostMealEntryRequest) Entry {
	return Entry{
		foodItemID: utils.NewNullInt64(r.FoodItemID, 0),
		recipeID:   utils.NewNullInt64(r.RecipeID, 0),
		Amount:     r.Amount,
	}
}
