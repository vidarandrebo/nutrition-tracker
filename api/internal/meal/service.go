package meal

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
)

type IService interface {
	Add(item *Meal) (*Meal, error)
	AddFoodItemEntry(entry *Entry[*fooditem.FoodItem], mealID int64, ownerID int64) (*Entry[*fooditem.FoodItem], error)
	AddMacroEntry(entry *Entry[*Macronutrient], mealID int64, ownerID int64) (*Entry[*Macronutrient], error)
	AddRecipeEntry(entry *Entry[*recipe.Recipe], mealID int64, ownerID int64) (*Entry[*recipe.Recipe], error)
}
