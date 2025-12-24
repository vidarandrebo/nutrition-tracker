package meal

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type RecipeEntry struct {
	ID             int64
	RecipeID       int64
	Amount         float64
	SequenceNumber int
}

func NewRecipeEntry() *RecipeEntry {
	return &RecipeEntry{}
}

func RMEFromRequest(r api.MealRecipeEntryPostRequest) *RecipeEntry {
	return &RecipeEntry{
		ID:             0,
		RecipeID:       r.RecipeId,
		Amount:         r.Amount,
		SequenceNumber: 0,
	}
}

func (rme *RecipeEntry) ToTable(mealID int64) TableMealRecipeEntry {
	return TableMealRecipeEntry{
		ID:             rme.ID,
		RecipeID:       rme.RecipeID,
		Amount:         rme.Amount,
		SequenceNumber: rme.SequenceNumber,
		MealID:         mealID,
	}
}

func (rme *RecipeEntry) FromTable(entry TableMealRecipeEntry) *RecipeEntry {
	rme.ID = entry.ID
	rme.RecipeID = entry.RecipeID
	rme.Amount = entry.Amount
	rme.SequenceNumber = entry.SequenceNumber
	return rme
}

func (rme *RecipeEntry) ToResponse() api.MealRecipeEntryResponse {
	return api.MealRecipeEntryResponse{
		Id:             rme.ID,
		Amount:         rme.Amount,
		RecipeId:       rme.RecipeID,
		SequenceNumber: rme.SequenceNumber,
	}
}
