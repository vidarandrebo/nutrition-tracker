package fooditem

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type Store struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewStore(db *sql.DB, logger *slog.Logger) *Store {
	s := Store{db: db}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}

func (s *Store) Add(item FoodItem) FoodItem {
	tx, err := s.db.Begin()
	tx.QueryRow(`
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
		panic(err)
	}
	for _, microNutrient := range item.Micronutrients {
		_, err = tx.Exec(`
			INSERT INTO micronutrients (name, amount, food_item_id) 
			VALUES ($1, $2, $3)`,
			microNutrient.Name, microNutrient.Amount, item.ID)
		if err != nil {
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	s.logger.Info("added fooditem", slog.Any("fooditem", item))

	return item
}

func (s *Store) GetByID(id int64) (FoodItem, error) {
	item := FoodItem{}
	err := s.db.QueryRow(`
		SELECT id, manufacturer, product, protein, carbohydrate, fat, kcal, public, source, owner_id 
		FROM food_items 
		WHERE id = $1`,
		id,
	).Scan(
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
		return FoodItem{}, err
	}
	return item, nil
}

func (s *Store) Get(ownerID int64) []FoodItem {
	items := make([]FoodItem, 0)
	rows, err := s.db.Query(`
		SELECT id, manufacturer, product, protein, carbohydrate, fat, kcal, public, source, owner_id 
		FROM food_items f
		WHERE f.public = TRUE 
		  OR f.owner_id = $1`,
		ownerID,
	)
	for rows.Next() {
		item := FoodItem{}
		rows.Scan(
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
		items = append(items, item)
	}
	if err != nil {
		panic(err)
	}
	return items
}
func (s *Store) Delete(id int64, ownerID int64) error {
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
