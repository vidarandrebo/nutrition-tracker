package meal

import "time"

type PostMealRequest struct {
	Timestamp time.Time `json:"timestamp"`
}

type PostMealEntryRequest struct {
	MealID     int64
	FoodItemID int64
	Amount     float64
	Timestamp  time.Time
}
