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
	GetFoodItemEntries(mealID int64, ownerID int64) ([]*FoodItemMealEntry, error)
	GetRecipeEntries(mealID int64, ownerID int64) ([]*RecipeMealEntry, error)
	GetMacronutrientEntries(mealID int64, ownerID int64) ([]*MacronutrientMealEntry, error)
	Delete(id int64, ownerID int64) error
}
type Service struct {
	logger     *slog.Logger
	repository IRepository
}

func (s Service) GetById(id int64, ownerID int64) (*Meal, error) {
	if err := s.repository.CheckOwnership(id, ownerID); err != nil {
		return nil, err
	}
	item, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	meal := FromMealTable(item)
	err = s.loadEntries(meal)
	if err != nil {
		return nil, err
	}
	return meal, nil
}

func (s Service) Delete(id int64, ownerID int64) error {
	if err := s.repository.CheckOwnership(id, ownerID); err != nil {
		return err
	}
	return s.repository.Delete(id)
}

func (s Service) AddFoodItemEntry(entry *FoodItemMealEntry, mealID int64, ownerID int64) (*FoodItemMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addFoodItemEntry(entry, mealID)
}

func (s Service) addFoodItemEntry(entry *FoodItemMealEntry, mealID int64) (*FoodItemMealEntry, error) {
	item, fieErr := s.repository.AddFoodItemEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, fieErr
	}
	entry.ID = item.ID
	return entry, nil
}

func (s Service) AddMacroEntry(entry *MacronutrientMealEntry, mealID int64, ownerID int64) (*MacronutrientMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addMacroEntry(entry, mealID)
}

func (s Service) addMacroEntry(entry *MacronutrientMealEntry, mealID int64) (*MacronutrientMealEntry, error) {
	item, fieErr := s.repository.AddMacronutrientEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, fieErr
	}
	entry.ID = item.ID
	return entry, nil
}

func (s Service) AddRecipeEntry(entry *RecipeMealEntry, mealID int64, ownerID int64) (*RecipeMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addRecipeEntry(entry, mealID)
}

func (s Service) addRecipeEntry(entry *RecipeMealEntry, mealID int64) (*RecipeMealEntry, error) {
	item, fieErr := s.repository.AddRecipeEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, fieErr
	}
	entry.ID = item.ID
	return entry, nil
}

func (s Service) Add(item *Meal) (*Meal, error) {
	meal, err := s.repository.Add(item.ToTable())
	if err != nil {
		return nil, err
	}

	item.ID = meal.ID

	for _, fie := range item.FoodItemEntries {
		_, fieErr := s.addFoodItemEntry(fie, item.ID)
		if fieErr != nil {
			return nil, fieErr
		}
	}

	for _, me := range item.MacronutrientEntries {
		_, meErr := s.addMacroEntry(me, item.ID)
		if meErr != nil {
			return nil, meErr
		}
	}

	for _, re := range item.RecipeEntries {
		_, reErr := s.addRecipeEntry(re, item.ID)
		if reErr != nil {
			return nil, reErr
		}
	}

	return item, nil
}

func (s Service) GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error) {
	items, err := s.repository.GetByDate(from, to, ownerID)
	if err != nil {
		return nil, err
	}

	meals := make([]*Meal, 0, len(items))
	for _, item := range items {
		meal := FromMealTable(item)
		loadErr := s.loadEntries(meal)
		if loadErr != nil {
			return nil, loadErr
		}
		meals = append(meals, meal)
	}
	return meals, nil
}

func (s Service) loadEntries(meal *Meal) error {
	foodItemEntries, err := s.getFoodItemEntries(meal.ID)
	if err != nil {
		return err
	}
	meal.FoodItemEntries = foodItemEntries

	recipeEntries, err := s.getRecipeEntries(meal.ID)
	if err != nil {
		return err
	}
	meal.RecipeEntries = recipeEntries

	macroEntries, err := s.getMacronutrientEntries(meal.ID)
	if err != nil {
		return err
	}
	meal.MacronutrientEntries = macroEntries

	return nil
}

func (s Service) GetFoodItemEntries(mealID int64, ownerID int64) ([]*FoodItemMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getFoodItemEntries(mealID)
}

func (s Service) getFoodItemEntries(mealID int64) ([]*FoodItemMealEntry, error) {
	items, err := s.repository.GetFoodItemEntries(mealID)
	if err != nil {
		return nil, err
	}
	entries := make([]*FoodItemMealEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, FromFoodItemMealEntryTable(item))
	}
	return entries, nil
}

func (s Service) GetRecipeEntries(mealID int64, ownerID int64) ([]*RecipeMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getRecipeEntries(mealID)
}

func (s Service) getRecipeEntries(mealID int64) ([]*RecipeMealEntry, error) {
	items, err := s.repository.GetRecipeEntries(mealID)
	if err != nil {
		return nil, err
	}
	entries := make([]*RecipeMealEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, FromRecipeMealEntryTable(item))
	}
	return entries, nil
}

func (s Service) GetMacronutrientEntries(mealID int64, ownerID int64) ([]*MacronutrientMealEntry, error) {
	if err := s.repository.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getMacronutrientEntries(mealID)
}

func (s Service) getMacronutrientEntries(mealID int64) ([]*MacronutrientMealEntry, error) {
	items, err := s.repository.GetMacronutrientEntries(mealID)
	if err != nil {
		return nil, err
	}
	entries := make([]*MacronutrientMealEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, FromMacronutrientMealEntryTable(item))
	}
	return entries, nil
}

func NewService(repository IRepository, logger *slog.Logger) *Service {
	s := Service{repository: repository}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}
