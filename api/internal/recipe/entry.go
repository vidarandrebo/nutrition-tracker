package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Entry struct {
	ID         int64
	Amount     float64
	FoodItemID int64
}

func (e *Entry) FoodItemIDOrNil() any {
	if e.FoodItemID == 0 {
		return nil
	}
	return e.FoodItemID
}

func (e *Entry) ToResponse() api.RecipeEntryResponse {
	return api.RecipeEntryResponse{
		Id:         e.ID,
		Amount:     e.Amount,
		FoodItemId: e.FoodItemID,
	}
}

func (e *Entry) ToTable(recipeID int64) TableRecipeFoodItemEntry {
	return TableRecipeFoodItemEntry{
		ID:         e.ID,
		Amount:     e.Amount,
		FoodItemID: e.FoodItemID,
		RecipeID:   recipeID,
	}
}

func FromRecipeEntryTable(tbl TableRecipeFoodItemEntry) *Entry {
	return &Entry{
		ID:         tbl.ID,
		Amount:     tbl.Amount,
		FoodItemID: tbl.FoodItemID,
	}
}

func FromEntryPost(r api.PostRecipeEntryRequest) *Entry {
	return &Entry{
		Amount:     r.Amount,
		FoodItemID: r.FoodItemId,
	}
}
