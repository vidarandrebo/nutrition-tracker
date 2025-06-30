package recipe

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type Entry struct {
	ID         int64
	Amount     float64
	FoodItemID int64
	RecipeID   int64
}

func (e Entry) FoodItemIDOrNil() any {
	if e.FoodItemID == 0 {
		return nil
	}
	return e.FoodItemID
}

func (e Entry) IsValid() bool {
	return (e.ID != 0) && (e.Amount != 0.0)
}

func (e Entry) ToResponse() api.RecipeEntryResponse {
	return api.RecipeEntryResponse{
		Id:         &e.ID,
		Amount:     &e.Amount,
		FoodItemId: &e.FoodItemID,
	}
}
