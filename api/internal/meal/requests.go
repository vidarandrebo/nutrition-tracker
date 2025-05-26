package meal

import "time"

type PostMealRequest struct {
	Timestamp time.Time `json:"timestamp"`
}

type PostMealEntryRequest struct {
	FoodItemID int64   `json:"foodItemId"`
	Amount     float64 `json:"amount"`
}
