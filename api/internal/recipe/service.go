package recipe

import (
	"log/slog"
	"reflect"
)

type IService interface {
	Get(ownerID int64) ([]*Recipe, error)
	GetById(id int64, ownerID int64) (*Recipe, error)
	Add(item *Recipe) (*Recipe, error)
	Delete(id int64, ownerID int64) error
}

type Service struct {
	repository IRepository
	logger     *slog.Logger
}

func NewService(repository IRepository, logger *slog.Logger) *Service {
	s := Service{repository: repository}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}

func (s *Service) Get(ownerID int64) ([]*Recipe, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		return nil, err
	}
	recipes := make([]*Recipe, 0, len(items))
	for _, item := range items {
		recipe := FromRecipeTable(item)
		recipeEntries, err := s.repository.GetEntries(item.ID)
		if err != nil {
			return nil, err
		}
		for _, re := range recipeEntries {
			recipe.FoodItemEntries = append(recipe.FoodItemEntries, FromRecipeEntryTable(re))
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (s *Service) GetById(id int64, ownerID int64) (*Recipe, error) {
	if err := s.repository.CheckOwnership(id, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	recipe := FromRecipeTable(item)
	recipeEntries, err := s.repository.GetEntries(item.ID)
	if err != nil {
		return nil, err
	}

	for _, re := range recipeEntries {
		recipe.FoodItemEntries = append(recipe.FoodItemEntries, FromRecipeEntryTable(re))
	}

	return recipe, nil
}

func (s *Service) Add(item *Recipe) (*Recipe, error) {
	recipe, rErr := s.repository.Add(item.ToTable())

	if rErr != nil {
		return nil, rErr
	}
	item.ID = recipe.ID

	for _, re := range item.FoodItemEntries {
		recipeEntry, reErr := s.repository.AddEntry(re.ToTable(item.ID))
		if reErr != nil {
			return nil, reErr
		}
		re.ID = recipeEntry.ID
	}
	return item, nil
}

func (s *Service) Delete(id int64, ownerID int64) error {
	if err := s.repository.CheckOwnership(id, ownerID); err != nil {
		return err
	}

	return s.repository.Delete(id)
}
