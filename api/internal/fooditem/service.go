package fooditem

import (
	"database/sql"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IService interface {
	Get(id int64) ([]*FoodItem, error)
	GetByID(id int64, ownerID int64) (*FoodItem, error)
	Add(item *FoodItem) (*FoodItem, error)
	Delete(item *FoodItem, ownerID int64) error
}

type Service struct {
	db         *sql.DB
	logger     *slog.Logger
	repository *Repository
}

func NewService(db *sql.DB, repository *Repository, logger *slog.Logger) *Service {
	s := Service{db: db, repository: repository}
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

func (s *Service) GetByID(id int64) (FoodItem, error) {
	rows, err := s.db.Query(`
		WITH owners_fi AS (
			SELECT * 
			FROM food_items
			WHERE id = $1
		)
		SELECT fi.id, fi.manufacturer, fi.product, fi.protein, fi.carbohydrate, fi.fat, fi.kcal, fi.public, fi.source, fi.owner_id, ps.id, ps.amount, ps.name, m.id, m.amount, m.name
		FROM owners_fi fi
		LEFT JOIN public.portion_sizes ps ON fi.id = ps.food_item_id
		LEFT JOIN micronutrients m ON fi.id = m.food_item_id
		`,
		id,
	)
	if err != nil {
		s.logger.Error("failed to get fooditem from database", slog.Any("err", err))
		return FoodItem{}, err
	}
	items := make([]TableFoodItemComplete, 0)
	for rows.Next() {
		item := TableFoodItemComplete{}
		rows.Scan(
			&item.FI.ID,
			&item.FI.Manufacturer,
			&item.FI.Product,
			&item.FI.Protein,
			&item.FI.Carbohydrate,
			&item.FI.Fat,
			&item.FI.KCal,
			&item.FI.Public,
			&item.FI.Source,
			&item.FI.OwnerID,
			&item.PS.ID,
			&item.PS.Amount,
			&item.PS.Name,
			&item.M.ID,
			&item.M.Amount,
			&item.M.Name,
		)
		items = append(items, item)
	}
	return fromFoodItemComplete(items)[0], nil
}

func (s *Service) Get(ownerID int64) ([]*FoodItem, error) {
	items, err := s.repository.Get(ownerID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) Delete(id int64, ownerID int64) error {
	_, err := s.db.Query(`
		DELETE FROM food_items
		WHERE id = $1
		  AND owner_id = $2
	`, id, ownerID,
	)
	if err != nil {
		s.logger.Error("failed to delete foodItem", slog.Int64("foodItemId", id))
		return err
	}
	return nil
}

func (s *Service) AddPortionSize(foodItemID int64, portionSize PortionSize, ownerID int64) (PortionSize, error) {
	if ok, err := s.ownsFoodItem(foodItemID, ownerID); !ok {
		if err != nil {
			return PortionSize{}, err
		}
		return PortionSize{}, err

	}
	err := s.db.QueryRow(`
			INSERT INTO portion_sizes AS ps (name, amount, food_item_id)
			VALUES ($1,$2,$3)
			RETURNING ps.id
		`, portionSize.Name, portionSize.Amount, foodItemID).Scan(&portionSize.ID)
	if err != nil {
		s.logger.Error("failed to add portionSize to foodItem", slog.Int64("foodItemId", foodItemID))
		return PortionSize{}, err
	}
	return portionSize, nil
}

func (s *Service) AddMicronutrient(foodItemID int64, micronutrient Micronutrient, ownerID int64) (Micronutrient, error) {
	if ok, err := s.ownsFoodItem(foodItemID, ownerID); !ok {
		if err != nil {
			return Micronutrient{}, err
		}
		return Micronutrient{}, err

	}
	err := s.db.QueryRow(`
			INSERT INTO micronutrients AS mn (name, amount, food_item_id)
			VALUES ($1,$2,$3)
			RETURNING mn.id
		`, micronutrient.Name, micronutrient.Amount, foodItemID).Scan(&micronutrient.ID)
	if err != nil {
		s.logger.Error("failed to add micronutrient to foodItem", slog.Int64("foodItemId", foodItemID))
		return Micronutrient{}, err
	}
	return micronutrient, nil
}

func (s *Service) ownsFoodItem(id int64, ownerID int64) (bool, error) {
	foodItem := FoodItem{}
	err := s.db.QueryRow(`
		SELECT id, owner_id 
		FROM food_items 
		WHERE id = $1
	`, id).Scan(
		&foodItem.ID,
		&foodItem.OwnerID,
	)
	if err != nil {
		return false, utils.ErrEntityNotFound
	}
	if foodItem.OwnerID == ownerID {
		return true, nil
	}
	return false, utils.ErrEntityNotOwned
}
