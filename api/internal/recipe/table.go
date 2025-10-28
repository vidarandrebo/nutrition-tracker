package recipe

import "time"

type TableRecipe struct {
	ID           int64
	Name         string
	DateCreated  time.Time
	DateModified time.Time
	OwnerID      int64
}

type TableRecipeEntry struct {
	ID           int64
	Amount       float64
	FoodItemID   int64
	DateCreated  time.Time
	DateModified time.Time
	RecipeID     int64
}
