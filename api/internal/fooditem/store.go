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
	for _, micronutrient := range item.Micronutrients {
		_, err = tx.Exec(`
			INSERT INTO micronutrients (name, amount, food_item_id) 
			VALUES ($1, $2, $3)`,
			micronutrient.Name, micronutrient.Amount, item.ID)
		if err != nil {
			panic(err)
		}
	}
	for _, portionSize := range item.PortionSizes {
		_, err = tx.Exec(`
			INSERT INTO portion_sizes (name, amount, food_item_id) 
			VALUES ($1, $2,$3)
    	`, portionSize.Name, portionSize.Amount, item.ID)
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
	rows, err := s.db.Query(`
		SELECT fi.id, fi.manufacturer, fi.product, fi.protein, fi.carbohydrate, fi.fat, fi.kcal, fi.public, fi.source, fi.owner_id, ps.id, ps.amount, ps.name
		FROM food_items fi
		JOIN public.portion_sizes ps ON fi.id = ps.food_item_id
		WHERE fi.id = $1`,
		id,
	)
	items := make([]TableFoodItemAndPortion, 0)
	for rows.Next() {
		item := TableFoodItemAndPortion{}
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
			&item.P.ID,
			&item.P.Amount,
			&item.P.Name,
		)
		items = append(items, item)
	}
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
