package meal

type Entry struct {
	ID         int64
	FoodItemID int64
	RecipeID   int64
	Amount     float64
}

func (e Entry) FoodItemIDOrNil() any {
	if e.FoodItemID == 0 {
		return nil
	}
	return e.FoodItemID
}

func (e Entry) ToResponse() EntryResponse {
	return EntryResponse{
		ID:         e.ID,
		Amount:     e.Amount,
		FoodItemID: e.FoodItemID,
	}
}

func (e Entry) IsValid() bool {
	return (e.ID != 0) && (e.Amount != 0.0)
}
