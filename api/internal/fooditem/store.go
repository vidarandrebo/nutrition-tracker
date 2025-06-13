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
		insert into food_items as fi (manufacturer, product, protein, carbohydrate, fat, kcal, source, owner_id) 
		values ($1, $2, $3, $4, $5, $6, $7, $8) 
		returning fi.id`,
		item.Manufacturer,
		item.Product,
		item.Protein,
		item.Carbohydrate,
		item.Fat,
		item.KCal,
		item.Source,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		panic(err)
	}
	for _, microNutrient := range item.Micronutrients {
		_, err = tx.Exec(`
			insert into micronutrients (name, amount, food_item_id) 
			values ($1, $2, $3)`,
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
		select id, manufacturer, product, protein, carbohydrate, fat, kcal,source, owner_id 
		from food_items 
		where id = $1`,
		id,
	).Scan(
		&item.ID,
		&item.Manufacturer,
		&item.Product,
		&item.Protein,
		&item.Carbohydrate,
		&item.Fat,
		&item.KCal,
		&item.Source,
		&item.OwnerID,
	)
	if err != nil {
		return FoodItem{}, err
	}
	return item, nil
}

func (s *Store) Get() []FoodItem {
	items := make([]FoodItem, 0)
	rows, err := s.db.Query(`
		select id, manufacturer, product, protein, carbohydrate, fat, kcal , source, owner_id 
		from food_items`)
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
