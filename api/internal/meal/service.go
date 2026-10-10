package meal

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IService interface {
	Add(item *Meal) (*Meal, error)
	AddFoodItemEntry(entry *FoodItemEntry, mealID int64, ownerID int64) (*FoodItemEntry, error)
	AddMacroEntry(entry *MacronutrientEntry, mealID int64, ownerID int64) (*MacronutrientEntry, error)
	AddRecipeEntry(entry *RecipeEntry, mealID int64, ownerID int64) (*RecipeEntry, error)
	GetById(id int64, ownerID int64) (*Meal, error)
	GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error)
	GetFoodItemEntries(mealID int64, ownerID int64) ([]*FoodItemEntry, error)
	GetRecipeEntries(mealID int64, ownerID int64) ([]*RecipeEntry, error)
	GetMacronutrientEntries(mealID int64, ownerID int64) ([]*MacronutrientEntry, error)
	Delete(id int64, ownerID int64) error
	DeleteFoodItemEntry(id int64, mealID int64, ownerID int64) error
	DeleteRecipeEntry(id int64, mealID int64, ownerID int64) error
	DeleteMacronutrientEntry(id int64, mealID, ownerID int64) error
	CheckOwnership(id int64, ownerID int64) error
}
type Service struct {
	logger     *slog.Logger
	repository IRepository
}

// repositoryError returns utils.ErrNotFound for sql.ErrNoRows or utils.ErrUnknown
// for other repository failures, logging the original error before translation.
func (s *Service) repositoryError(operation string, id int64, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		s.logger.Debug("meal not found", slog.String("operation", operation), slog.Int64("mealID", id), slog.Any("err", err))
		return utils.ErrNotFound
	}
	s.logger.Error("meal repository failure", slog.String("operation", operation), slog.Int64("mealID", id), slog.Any("err", err))
	return utils.ErrUnknown
}

// CheckOwnership returns utils.ErrNotFound when the meal does not exist,
// utils.ErrUnknown for other lookup errors, or utils.ErrForbidden when it has another owner.
func (s *Service) CheckOwnership(id int64, ownerID int64) error {
	item, err := s.repository.GetById(id)
	if err != nil {
		return s.repositoryError("check ownership", id, err)
	}
	if item.OwnerID != ownerID {
		return utils.ErrForbidden
	}
	return nil
}

func (s *Service) GetById(id int64, ownerID int64) (*Meal, error) {
	if err := s.CheckOwnership(id, ownerID); err != nil {
		return nil, err
	}
	item, err := s.repository.GetById(id)
	if err != nil {
		return nil, s.repositoryError("get by ID", id, err)
	}
	meal := NewMeal().FromTable(item)
	err = s.loadEntries(meal)
	if err != nil {
		return nil, err
	}
	return meal, nil
}

func (s *Service) Delete(id int64, ownerID int64) error {
	if err := s.CheckOwnership(id, ownerID); err != nil {
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		return s.repositoryError("delete", id, err)
	}
	return nil
}

func (s *Service) AddFoodItemEntry(entry *FoodItemEntry, mealID int64, ownerID int64) (*FoodItemEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addFoodItemEntry(entry, mealID)
}

func (s *Service) addFoodItemEntry(entry *FoodItemEntry, mealID int64) (*FoodItemEntry, error) {
	item, fieErr := s.repository.AddFoodItemEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, s.repositoryError("add food item entry", mealID, fieErr)
	}
	entry.ID = item.ID
	return entry, nil
}

func (s *Service) AddMacroEntry(entry *MacronutrientEntry, mealID int64, ownerID int64) (*MacronutrientEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addMacroEntry(entry, mealID)
}

func (s *Service) addMacroEntry(entry *MacronutrientEntry, mealID int64) (*MacronutrientEntry, error) {
	item, fieErr := s.repository.AddMacronutrientEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, s.repositoryError("add macronutrient entry", mealID, fieErr)
	}
	entry.ID = item.ID
	return entry, nil
}

func (s *Service) AddRecipeEntry(entry *RecipeEntry, mealID int64, ownerID int64) (*RecipeEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.addRecipeEntry(entry, mealID)
}

func (s *Service) addRecipeEntry(entry *RecipeEntry, mealID int64) (*RecipeEntry, error) {
	item, fieErr := s.repository.AddRecipeEntry(entry.ToTable(mealID))
	if fieErr != nil {
		return nil, s.repositoryError("add recipe entry", mealID, fieErr)
	}
	entry.ID = item.ID
	return entry, nil
}

func (s *Service) Add(item *Meal) (*Meal, error) {
	meal, err := s.repository.Add(item.ToTable())
	if err != nil {
		return nil, s.repositoryError("add", item.OwnerID, err)
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

func (s *Service) GetByDate(from time.Time, to time.Time, ownerID int64) ([]*Meal, error) {
	items, err := s.repository.GetByDate(from, to, ownerID)
	if err != nil {
		return nil, s.repositoryError("list by date", ownerID, err)
	}

	meals := make([]*Meal, 0, len(items))
	for _, item := range items {
		meal := NewMeal().FromTable(item)
		loadErr := s.loadEntries(meal)
		if loadErr != nil {
			return nil, loadErr
		}
		meals = append(meals, meal)
	}
	return meals, nil
}

func (s *Service) loadEntries(meal *Meal) error {
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

func (s *Service) GetFoodItemEntries(mealID int64, ownerID int64) ([]*FoodItemEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getFoodItemEntries(mealID)
}

func (s *Service) getFoodItemEntries(mealID int64) ([]*FoodItemEntry, error) {
	items, err := s.repository.GetFoodItemEntries(mealID)
	if err != nil {
		return nil, s.repositoryError("get food item entries", mealID, err)
	}
	entries := make([]*FoodItemEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, NewFoodItemEntry().FromTable(item))
	}
	return entries, nil
}

func (s *Service) GetRecipeEntries(mealID int64, ownerID int64) ([]*RecipeEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getRecipeEntries(mealID)
}

func (s *Service) getRecipeEntries(mealID int64) ([]*RecipeEntry, error) {
	items, err := s.repository.GetRecipeEntries(mealID)
	if err != nil {
		return nil, s.repositoryError("get recipe entries", mealID, err)
	}
	entries := make([]*RecipeEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, NewRecipeEntry().FromTable(item))
	}
	return entries, nil
}

func (s *Service) GetMacronutrientEntries(mealID int64, ownerID int64) ([]*MacronutrientEntry, error) {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return nil, err
	}
	return s.getMacronutrientEntries(mealID)
}

func (s *Service) getMacronutrientEntries(mealID int64) ([]*MacronutrientEntry, error) {
	items, err := s.repository.GetMacronutrientEntries(mealID)
	if err != nil {
		return nil, s.repositoryError("get macronutrient entries", mealID, err)
	}
	entries := make([]*MacronutrientEntry, 0, len(items))
	for _, item := range items {
		entries = append(entries, NewMacronutrientEntry().FromTable(item))
	}
	return entries, nil
}

func (s *Service) DeleteFoodItemEntry(id int64, mealID int64, ownerID int64) error {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return err
	}
	if err := s.repository.DeleteFoodItemEntry(id, mealID); err != nil {
		return s.repositoryError("delete food item entry", mealID, err)
	}
	return nil
}

func (s *Service) DeleteRecipeEntry(id int64, mealID int64, ownerID int64) error {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return err
	}
	if err := s.repository.DeleteRecipeEntry(id, mealID); err != nil {
		return s.repositoryError("delete recipe entry", mealID, err)
	}
	return nil
}

func (s *Service) DeleteMacronutrientEntry(id int64, mealID int64, ownerID int64) error {
	if err := s.CheckOwnership(mealID, ownerID); err != nil {
		return err
	}
	if err := s.repository.DeleteMacronutrientEntry(id, mealID); err != nil {
		return s.repositoryError("delete macronutrient entry", mealID, err)
	}
	return nil
}

func NewService(repository IRepository, logger *slog.Logger) *Service {
	s := Service{repository: repository}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}
