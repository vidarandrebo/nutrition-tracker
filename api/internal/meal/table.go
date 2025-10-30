package meal

import "time"

type TableMeal struct {
	ID             int64
	SequenceNumber int
	MealTime       time.Time
	DateCreated    time.Time
	DateModified   time.Time
	OwnerID        int64
}
type TableMealEntry struct {
	ID           int64
	Amount       float64
	FoodItemID   int64
	RecipeID     int64
	DateCreated  time.Time
	DateModified time.Time
	MealID       int64
}
