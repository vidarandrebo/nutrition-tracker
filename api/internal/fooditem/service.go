package fooditem

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IService interface {
	Get(ownerID int64) ([]*FoodItem, error)
	GetByID(id int64) (*FoodItem, error)
	Add(item *FoodItem) (*FoodItem, error)
	AddPortionSize(item *PortionSize, foodItemID int64, ownerID int64) (*PortionSize, error)
	AddMicronutrient(item *Micronutrient, foodItemID int64, ownerID int64) (*Micronutrient, error)
	Delete(id int64, ownerID int64) error
	CheckOwnership(id int64, ownerID int64) error
}

type Service struct {
	logger     *slog.Logger
	repository IRepository
}

func NewService(repository IRepository, logger *slog.Logger) *Service {
	s := Service{repository: repository}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}

// Add returns utils.ErrUnknown for any food item, micronutrient, or portion size insert failure.
func (s *Service) Add(item *FoodItem) (*FoodItem, error) {
	foodItem, fiErr := s.repository.Add(item.ToTable())

	if fiErr != nil {
		s.logger.Error("failed to add food item", slog.Int64("ownerID", item.OwnerID), slog.Any("err", fiErr))
		return nil, utils.ErrUnknown
	}

	item.ID = foodItem.ID

	for _, mn := range item.Micronutrients {
		micronutrient, mnErr := s.repository.AddMicronutrient(mn.ToTable(foodItem.ID))
		if mnErr != nil {
			s.logger.Error("failed to add micronutrient", slog.Int64("foodItemID", foodItem.ID), slog.Any("err", mnErr))
			return nil, utils.ErrUnknown
		}
		mn.ID = micronutrient.ID
	}
	for _, ps := range item.PortionSizes {
		portionSize, psErr := s.repository.AddPortionSize(ps.ToTable(item.ID))
		if psErr != nil {
			s.logger.Error("failed to add portion size", slog.Int64("foodItemID", foodItem.ID), slog.Any("err", psErr))
			return nil, utils.ErrUnknown
		}
		ps.ID = portionSize.ID
	}
	return item, nil
}

// GetByID returns utils.ErrNotFound if the food item is missing, or utils.ErrUnknown
// for other food item lookup or portion size list failures.
func (s *Service) GetByID(id int64) (*FoodItem, error) {
	item, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Debug("food item not found", slog.Int64("foodItemID", id), slog.Any("err", err))
			return nil, utils.ErrNotFound
		}
		s.logger.Error("failed to get food item", slog.Int64("foodItemID", id), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	foodItem := NewFoodItem().FromTable(item)
	portionSizes, err := s.repository.GetPortionSizes(id)
	if err != nil {
		s.logger.Error("failed to get portion sizes", slog.Int64("foodItemID", id), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	for _, ps := range portionSizes {
		foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
	}
	return foodItem, nil
}

// Get returns utils.ErrUnknown for food item or portion size list failures.
// An owner with no food items receives an empty result without error.
func (s *Service) Get(ownerID int64) ([]*FoodItem, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		s.logger.Error("failed to list food items", slog.Int64("ownerID", ownerID), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	foodItems := make([]*FoodItem, 0, len(items))
	for _, item := range items {
		foodItem := NewFoodItem().FromTable(item)
		portionSizes, err := s.repository.GetPortionSizes(item.ID)
		if err != nil {
			s.logger.Error("failed to get portion sizes", slog.Int64("foodItemID", item.ID), slog.Any("err", err))
			return nil, utils.ErrUnknown
		}
		for _, ps := range portionSizes {
			foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
		}
		foodItems = append(foodItems, foodItem)
	}

	return foodItems, nil
}

// CheckOwnership returns utils.ErrNotFound for a missing food item,
// utils.ErrUnknown for lookup failures, or utils.ErrForbidden for another owner.
func (s *Service) CheckOwnership(id int64, ownerID int64) error {
	item, err := s.repository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Debug("food item not found", slog.Int64("foodItemID", id), slog.Any("err", err))
			return utils.ErrNotFound
		}
		s.logger.Error("failed to check food item ownership", slog.Int64("foodItemID", id), slog.Any("err", err))
		return utils.ErrUnknown
	}
	if item.OwnerID != ownerID {
		return utils.ErrForbidden
	}
	return nil
}

// AddPortionSize returns utils.ErrNotFound for a missing food item,
// utils.ErrForbidden for another owner, or utils.ErrUnknown for lookup or insert failures.
func (s *Service) AddPortionSize(portionSize *PortionSize, foodItemID int64, ownerID int64) (*PortionSize, error) {
	if err := s.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddPortionSize(portionSize.ToTable(foodItemID))
	if err != nil {
		s.logger.Error("failed to add portion size", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	portionSize.ID = item.ID
	return portionSize, nil
}

// AddMicronutrient returns utils.ErrNotFound for a missing food item,
// utils.ErrForbidden for another owner, or utils.ErrUnknown for lookup or insert failures.
func (s *Service) AddMicronutrient(micronutrient *Micronutrient, foodItemID int64, ownerID int64) (*Micronutrient, error) {
	if err := s.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddMicronutrient(micronutrient.ToTable(foodItemID))
	if err != nil {
		s.logger.Error("failed to add micronutrient", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	micronutrient.ID = item.ID
	return micronutrient, nil
}

// Delete returns utils.ErrNotFound for a missing food item,
// utils.ErrForbidden for another owner, or utils.ErrUnknown for lookup or delete failures.
func (s *Service) Delete(id int64, ownerID int64) error {
	if err := s.CheckOwnership(id, ownerID); err != nil {
		return err
	}
	if err := s.repository.Delete(id); err != nil {
		s.logger.Error("failed to delete food item", slog.Int64("foodItemID", id), slog.Any("err", err))
		return utils.ErrUnknown
	}
	return nil
}
