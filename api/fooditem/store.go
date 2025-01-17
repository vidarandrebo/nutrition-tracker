package fooditem

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetFoodItem() *FoodItem {
	item := FoodItem{}
	err := s.db.QueryRow("select id, name from food_items").Scan(&item.ID, &item.Name)
	if err != nil {
		panic(err)
	}

	return &item
}
