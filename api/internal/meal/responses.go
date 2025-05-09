package meal

import "time"

type MealResponse struct {
	ID             int64           `json:"id"`
	SequenceNumber int64           `json:"sequenceNumber"`
	Timestamp      time.Time       `json:"timestamp"`
	Entries        []EntryResponse `json:"entries"`
}

type EntryResponse struct {
	ID         int64   `json:"id"`
	Amount     float64 `json:"amount"`
	FoodItemID int64   `json:"foodItemID"`
}
