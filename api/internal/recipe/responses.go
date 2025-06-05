package recipe

type RecipeResponse struct {
	ID      int64           `json:"id"`
	Name    string          `json:"name"`
	Entries []EntryResponse `json:"entries"`
}

type EntryResponse struct {
	ID         int64   `json:"id"`
	Amount     float64 `json:"amount"`
	FoodItemID int64   `json:"foodItemId"`
}
