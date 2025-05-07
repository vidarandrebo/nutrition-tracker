package meal

type MealResponse struct {
	ID      int64           `json:"id"`
	Entries []EntryResponse `json:"entries"`
}

type EntryResponse struct {
	ID         int64   `json:"id"`
	Amount     float64 `json:"amount"`
	FoodItemID int64   `json:"foodItemID"`
}
