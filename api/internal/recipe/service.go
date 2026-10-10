package recipe

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
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

// repositoryError returns utils.ErrNotFound for sql.ErrNoRows or utils.ErrUnknown
// for other repository failures, logging the original error before translation.
func (s *Service) repositoryError(operation string, id int64, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		s.logger.Debug("recipe not found", slog.String("operation", operation), slog.Int64("recipeID", id), slog.Any("err", err))
		return utils.ErrNotFound
	}
	s.logger.Error("recipe repository failure", slog.String("operation", operation), slog.Int64("recipeID", id), slog.Any("err", err))
	return utils.ErrUnknown
}

// CheckOwnership returns utils.ErrNotFound for a missing recipe,
// utils.ErrUnknown for lookup failures, or utils.ErrForbidden for another owner.
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

func (s *Service) Get(ownerID int64) ([]*Recipe, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		return nil, s.repositoryError("list", ownerID, err)
	}
	recipes := make([]*Recipe, 0, len(items))
	for _, item := range items {
		recipe := FromRecipeTable(item)
		recipeEntries, err := s.repository.GetEntries(item.ID)
		if err != nil {
			return nil, s.repositoryError("get entries", item.ID, err)
		}
		for _, re := range recipeEntries {
			recipe.FoodItemEntries = append(recipe.FoodItemEntries, FromRecipeEntryTable(re))
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (s *Service) GetById(id int64, ownerID int64) (*Recipe, error) {
	if err := s.CheckOwnership(id, ownerID); err != nil {
		return nil, err
	}

	item, err := s.repository.GetByID(id)
	if err != nil {
		return nil, s.repositoryError("get by ID", id, err)
	}

	recipe := FromRecipeTable(item)
	recipeEntries, err := s.repository.GetEntries(item.ID)
	if err != nil {
		return nil, s.repositoryError("get entries", id, err)
	}

	for _, re := range recipeEntries {
		recipe.FoodItemEntries = append(recipe.FoodItemEntries, FromRecipeEntryTable(re))
	}

	return recipe, nil
}

func (s *Service) Add(item *Recipe) (*Recipe, error) {
	recipe, rErr := s.repository.Add(item.ToTable())

	if rErr != nil {
		return nil, s.repositoryError("add", item.OwnerID, rErr)
	}
	item.ID = recipe.ID

	for _, re := range item.FoodItemEntries {
		recipeEntry, reErr := s.repository.AddEntry(re.ToTable(item.ID))
		if reErr != nil {
			return nil, s.repositoryError("add entry", item.ID, reErr)
		}
		re.ID = recipeEntry.ID
	}
	return item, nil
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
