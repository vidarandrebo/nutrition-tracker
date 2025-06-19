package recipe

type PostRecipeRequest struct {
	Name    string                   `json:"name"`
	Entries []PostRecipeEntryRequest `json:"entries"`
}
type PostRecipeEntryRequest struct {
	Amount     float64 `json:"amount"`
	FoodItemID int64   `json:"foodItemId"`
}
