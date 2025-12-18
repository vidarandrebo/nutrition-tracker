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

type TableMealRecipeEntry struct {
	ID             int64
	RecipeID       int64
	Amount         float64
	SequenceNumber int
	DateCreated    time.Time
	DateModified   time.Time
	MealID         int64
}

type TableMealFoodItemEntry struct {
	ID             int64
	FoodItemID     int64
	Amount         float64
	SequenceNumber int
	DateCreated    time.Time
	DateModified   time.Time
	MealID         int64
}

type TableMealMacronutrientEntry struct {
	ID             int64
	SequenceNumber int
	Protein        float64
	Carbohydrate   float64
	Fat            float64
	KCal           float64
	DateCreated    time.Time
	DateModified   time.Time
	MealID         int64
}
