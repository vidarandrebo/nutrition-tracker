package meal

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type FoodItemEntry struct {
	ID             int64
	FoodItemID     int64
	Amount         float64
	SequenceNumber int
}

func NewFoodItemEntry() *FoodItemEntry {
	return &FoodItemEntry{}
}

func FIMEFromRequest(r api.MealFoodItemEntryPostRequest) *FoodItemEntry {
	return &FoodItemEntry{
		ID:             0,
		FoodItemID:     r.FoodItemId,
		Amount:         r.Amount,
		SequenceNumber: 0,
	}
}

func (fime *FoodItemEntry) ToTable(mealID int64) TableMealFoodItemEntry {
	return TableMealFoodItemEntry{
		ID:             fime.ID,
		FoodItemID:     fime.FoodItemID,
		Amount:         fime.Amount,
		SequenceNumber: fime.SequenceNumber,
		MealID:         mealID,
	}
}

func (fime *FoodItemEntry) FromTable(entry TableMealFoodItemEntry) *FoodItemEntry {
	fime.ID = entry.ID
	fime.FoodItemID = entry.FoodItemID
	fime.Amount = entry.Amount
	fime.SequenceNumber = entry.SequenceNumber
	return fime
}

func (fime *FoodItemEntry) ToResponse() api.MealFoodItemEntryResponse {
	return api.MealFoodItemEntryResponse{
		Id:             fime.ID,
		Amount:         fime.Amount,
		FoodItemId:     fime.FoodItemID,
		SequenceNumber: fime.SequenceNumber,
	}
}
