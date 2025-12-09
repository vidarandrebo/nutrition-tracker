package meal

import (
	"log/slog"
	"reflect"
	"time"
)

type IService interface {
	Add(item *Meal) (*Meal, error)
	AddFoodItemEntry(entry *FoodItemMealEntry, mealID int64, ownerID int64) (*FoodItemMealEntry, error)
	AddMacroEntry(entry *MacronutrientMealEntry, mealID int64, ownerID int64) (*MacronutrientMealEntry, error)
	AddRecipeEntry(entry *RecipeMealEntry, mealID int64, ownerID int64) (*RecipeMealEntry, error)
	GetById(id int64, ownerID int64) (*Meal, error)
	GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error)
	Delete(id int64, ownerID int64) error
}
type Service struct {
	logger     *slog.Logger
	repository IRepository
}

func (s Service) AddFoodItemEntry(entry *FoodItemMealEntry, mealID int64, ownerID int64) (*FoodItemMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) AddMacroEntry(entry *MacronutrientMealEntry, mealID int64, ownerID int64) (*MacronutrientMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) AddRecipeEntry(entry *RecipeMealEntry, mealID int64, ownerID int64) (*RecipeMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Add(item *Meal) (*Meal, error) {
	meal, err := s.repository.Add(item.ToTable())

	if err != nil {
		return nil, err
	}

	item.ID = meal.ID

	for _, fie := range item.FoodItemEntries {
		foodItemEntry, fieErr := s.repository.AddFoodItemEntry(fie.ToTable(item.ID))
		if fieErr != nil {
			return nil, fieErr
		}
		fie.ID = foodItemEntry.ID
	}

	for _, me := range item.MacronutrientEntries {
		macronutrientEntry, meErr := s.repository.AddMacronutrientEntry(me.ToTable(item.ID))
		if meErr != nil {
			return nil, meErr
		}
		me.ID = macronutrientEntry.ID
	}

	for _, re := range item.RecipeEntries {
		recipeEntry, reErr := s.repository.AddRecipeEntry(re.ToTable(item.ID))
		if reErr != nil {
			return nil, reErr
		}
		re.ID = recipeEntry.ID
	}

	return item, nil
}

func (s Service) GetById(id int64, ownerID int64) (*Meal, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(id int64, ownerID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repository IRepository, logger *slog.Logger) *Service {
	s := Service{repository: repository}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}
