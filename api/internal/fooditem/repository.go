package fooditem

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type Repository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	r := Repository{db: db}
	r.logger = logger.With(slog.Any("module", reflect.TypeOf(r)))
	return &r
}

func (s *Repository) Add(item TableFoodItem) (TableFoodItem, error) {
	err := s.db.QueryRow(`
		INSERT INTO food_items AS fi (manufacturer, product, protein, carbohydrate, fat, kcal, public, source, owner_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING fi.id`,
		item.Manufacturer,
		item.Product,
		item.Protein,
		item.Carbohydrate,
		item.Fat,
		item.KCal,
		item.Public,
		item.Source,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		s.logger.Error("failed to add food item", slog.Int64("userID", item.OwnerID), slog.Any("err", err))
		return TableFoodItem{}, err
	}
	return item, nil
}

func (s *Repository) AddMicronutrient(item TableMicronutrient) (TableMicronutrient, error) {
	err := s.db.QueryRow(`
			INSERT INTO micronutrients (name, amount, food_item_id) 
			VALUES ($1, $2, $3)`,
		item.Name,
		item.Amount,
		item.FoodItemID,
	).Scan(&item.ID)

	if err != nil {
		s.logger.Error("failed to add micronutrient to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TableMicronutrient{}, err
	}

	return item, nil
}

func (s *Repository) AddPortionSize(item TablePortionSize) (TablePortionSize, error) {
	err := s.db.QueryRow(`
			INSERT INTO portion_sizes (name, amount, food_item_id) 
			VALUES ($1, $2,$3)
    	`, item.Name, item.Amount, item.ID,
	).Scan(&item.ID)

	if err != nil {
		s.logger.Error("failed to add portion size to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TablePortionSize{}, err
	}
	return item, nil
}

func (s *Repository) Get(ownerID int64) ([]TableFoodItem, error) {
	items := make([]TableFoodItem, 0)
	rows, err := s.db.Query(`
		WITH owned_fi AS (
		    SELECT *
			FROM food_items
			WHERE public = TRUE 
 			  OR owner_id = $1
        )		
		SELECT fi.id, fi.manufacturer, fi.product, fi.protein, fi.carbohydrate, fi.fat, fi.kcal, fi.public, fi.source, fi.owner_id
		FROM owned_fi fi
		`,
		ownerID,
	)
	if err != nil {
		s.logger.Error("failed to query rows of food items", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TableFoodItem{}
		err = rows.Scan(
			&item.ID,
			&item.Manufacturer,
			&item.Product,
			&item.Protein,
			&item.Carbohydrate,
			&item.Fat,
			&item.KCal,
			&item.Public,
			&item.Source,
			&item.OwnerID,
		)
		if err != nil {
			s.logger.Error("failed to scan rows of food items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Repository) GetPortionSizes(foodItemID int64) ([]TablePortionSize, error) {
	items := make([]TablePortionSize, 0)
	rows, err := s.db.Query(`
		WITH fi_portions AS (
		    SELECT *
			FROM portion_sizes
			WHERE food_item_id = $1
        )		
		SELECT ps.id, ps.name, ps.amount, ps.food_item_id
		FROM fi_portions ps
		`,
		foodItemID,
	)
	if err != nil {
		s.logger.Error("failed to query rows of portion sizes", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TablePortionSize{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			s.logger.Error("failed to scan rows of portion sizes items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Repository) GetMicronutrients(foodItemID int64) ([]TableMicronutrient, error) {
	items := make([]TableMicronutrient, 0)
	rows, err := s.db.Query(`
		WITH fi_micronutrients AS (
		    SELECT *
			FROM micronutrients
			WHERE food_item_id = $1
        )		
		SELECT mn.id, mn.name, mn.amount, mn.food_item_id
		FROM fi_micronutrients mn
		`,
		foodItemID,
	)
	if err != nil {
		s.logger.Error("failed to query rows of micronutrients", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TableMicronutrient{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			s.logger.Error("failed to scan rows of micronutrient items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
