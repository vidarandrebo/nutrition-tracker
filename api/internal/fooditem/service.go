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

// repositoryError returns utils.ErrNotFound for sql.ErrNoRows or utils.ErrUnknown
// for other repository failures, logging the original error before translation.
func (s *Service) repositoryError(operation string, id int64, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		s.logger.Debug("food item not found", slog.String("operation", operation), slog.Int64("foodItemID", id), slog.Any("err", err))
		return utils.ErrNotFound
	}
	s.logger.Error("food item repository failure", slog.String("operation", operation), slog.Int64("foodItemID", id), slog.Any("err", err))
	return utils.ErrUnknown
}

func (s *Service) Add(item *FoodItem) (*FoodItem, error) {
	foodItem, fiErr := s.repository.Add(item.ToTable())

	if fiErr != nil {
		return nil, s.repositoryError("add", item.OwnerID, fiErr)
	}

	item.ID = foodItem.ID

	for _, mn := range item.Micronutrients {
		micronutrient, mnErr := s.repository.AddMicronutrient(mn.ToTable(foodItem.ID))
		if mnErr != nil {
			return nil, s.repositoryError("add micronutrient", foodItem.ID, mnErr)
		}
		mn.ID = micronutrient.ID
	}
	for _, ps := range item.PortionSizes {
		portionSize, psErr := s.repository.AddPortionSize(ps.ToTable(item.ID))
		if psErr != nil {
			return nil, s.repositoryError("add portion size", foodItem.ID, psErr)
		}
		ps.ID = portionSize.ID
	}
	return item, nil
}

func (s *Service) GetByID(id int64) (*FoodItem, error) {
	item, err := s.repository.GetByID(id)
	if err != nil {
		return nil, s.repositoryError("get by ID", id, err)
	}
	foodItem := NewFoodItem().FromTable(item)
	portionSizes, err := s.repository.GetPortionSizes(id)
	if err != nil {
		return nil, s.repositoryError("get portion sizes", id, err)
	}
	for _, ps := range portionSizes {
		foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
	}
	return foodItem, nil
}

func (s *Service) Get(ownerID int64) ([]*FoodItem, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		return nil, s.repositoryError("list", ownerID, err)
	}
	foodItems := make([]*FoodItem, 0, len(items))
	for _, item := range items {
		foodItem := NewFoodItem().FromTable(item)
		portionSizes, err := s.repository.GetPortionSizes(item.ID)
		if err != nil {
			return nil, s.repositoryError("get portion sizes", item.ID, err)
		}
		for _, ps := range portionSizes {
			foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
		}
		foodItems = append(foodItems, foodItem)
	}

	return foodItems, nil
}

func (s *Service) CheckOwnership(id int64, ownerID int64) error {
	item, err := s.repository.GetByID(id)
	if err != nil {
		return s.repositoryError("check ownership", id, err)
	}
	if item.OwnerID != ownerID {
		return utils.ErrForbidden
	}
	return nil
}

func (s *Service) AddPortionSize(portionSize *PortionSize, foodItemID int64, ownerID int64) (*PortionSize, error) {
	if err := s.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddPortionSize(portionSize.ToTable(foodItemID))
	if err != nil {
		return nil, s.repositoryError("add portion size", foodItemID, err)
	}
	portionSize.ID = item.ID
	return portionSize, nil
}

func (s *Service) AddMicronutrient(micronutrient *Micronutrient, foodItemID int64, ownerID int64) (*Micronutrient, error) {
	if err := s.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddMicronutrient(micronutrient.ToTable(foodItemID))
	if err != nil {
		return nil, s.repositoryError("add micronutrient", foodItemID, err)
	}
	micronutrient.ID = item.ID
	return micronutrient, nil
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
