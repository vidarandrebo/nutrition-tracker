package fooditem

import (
	"log/slog"
	"reflect"
)

type IService interface {
	Get(ownerID int64) ([]*FoodItem, error)
	GetByID(id int64) (*FoodItem, error)
	Add(item *FoodItem) (*FoodItem, error)
	AddPortionSize(item *PortionSize, foodItemID int64, ownerID int64) (*PortionSize, error)
	AddMicronutrient(item *Micronutrient, foodItemID int64, ownerID int64) (*Micronutrient, error)
	Delete(id int64, ownerID int64) error
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

func (s *Service) Add(item *FoodItem) (*FoodItem, error) {
	foodItem, fiErr := s.repository.Add(item.ToTable())

	if fiErr != nil {
		return nil, fiErr
	}

	item.ID = foodItem.ID

	for _, mn := range item.Micronutrients {
		micronutrient, mnErr := s.repository.AddMicronutrient(mn.ToTable(foodItem.ID))
		if mnErr != nil {
			return nil, mnErr
		}
		mn.ID = micronutrient.ID
	}
	for _, ps := range item.PortionSizes {
		portionSize, psErr := s.repository.AddPortionSize(ps.ToTable(item.ID))
		if psErr != nil {
			return nil, psErr
		}
		ps.ID = portionSize.ID
	}
	return item, nil
}

func (s *Service) GetByID(id int64) (*FoodItem, error) {
	item, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	foodItem := NewFoodItem().FromTable(item)
	portionSizes, err := s.repository.GetPortionSizes(id)
	if err != nil {
		return nil, err
	}
	for _, ps := range portionSizes {
		foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
	}
	return foodItem, nil
}

func (s *Service) Get(ownerID int64) ([]*FoodItem, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		return nil, err
	}
	foodItems := make([]*FoodItem, 0, len(items))
	for _, item := range items {
		foodItem := NewFoodItem().FromTable(item)
		portionSizes, err := s.repository.GetPortionSizes(item.ID)
		if err != nil {
			return nil, err
		}
		for _, ps := range portionSizes {
			foodItem.PortionSizes = append(foodItem.PortionSizes, FromPortionSizeTable(ps))
		}
		foodItems = append(foodItems, foodItem)
	}

	return foodItems, nil
}

func (s *Service) AddPortionSize(portionSize *PortionSize, foodItemID int64, ownerID int64) (*PortionSize, error) {
	if err := s.repository.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddPortionSize(portionSize.ToTable(foodItemID))
	if err != nil {
		return nil, err
	}
	portionSize.ID = item.ID
	return portionSize, nil
}

func (s *Service) AddMicronutrient(micronutrient *Micronutrient, foodItemID int64, ownerID int64) (*Micronutrient, error) {
	if err := s.repository.CheckOwnership(foodItemID, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.AddMicronutrient(micronutrient.ToTable(foodItemID))
	if err != nil {
		return nil, err
	}
	micronutrient.ID = item.ID
	return micronutrient, nil
}

func (s *Service) Delete(id int64, ownerID int64) error {
	if err := s.repository.CheckOwnership(id, ownerID); err != nil {
		return err
	}
	return s.repository.Delete(id)
}
