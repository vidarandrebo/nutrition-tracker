package meal

import (
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
)

type IService interface {
	Add(item *Meal) (*Meal, error)
	AddFoodItemEntry(entry *Entry[*fooditem.FoodItem], mealID int64, ownerID int64) (*Entry[*fooditem.FoodItem], error)
	AddMacroEntry(entry *Entry[*Macronutrient], mealID int64, ownerID int64) (*Entry[*Macronutrient], error)
	AddRecipeEntry(entry *Entry[*recipe.Recipe], mealID int64, ownerID int64) (*Entry[*recipe.Recipe], error)
	GetById(id int64, ownerID int64) (*Meal, error)
	GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error)
	Delete(id int64, ownerID int64) error
}
