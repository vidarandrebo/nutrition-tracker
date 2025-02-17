package fooditem

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddFoodItem(request *PostFoodItemRequest, userID int64) *FoodItem {
	item := request.ToFoodItem()
	item.OwnerID = userID

	err := s.db.QueryRow("insert into food_items as fi (manufacturer, product, protein, carbohydrate, fat, kcal) values ($1, $2, $3, $4, $5, $6, $7) returning fi.id",
		item.Manufacturer,
		item.Product,
		item.Protein,
		item.Carbohydrate,
		item.Fat,
		item.KCal,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		panic(err)
	}
	return &item
}
func (s *Store) GetFoodItem(id int64) *FoodItem {
	item := FoodItem{}
	err := s.db.QueryRow("select id, manufacturer, product, protein, carbohydrate, fat, kcal, owner_id from food_items where id = $1", id).Scan(
		&item.ID,
		&item.Manufacturer,
		&item.Product,
		&item.Protein,
		&item.Carbohydrate,
		&item.Fat,
		&item.KCal,
		&item.OwnerID,
	)
	if err != nil {
		panic(err)
	}
	return &item
}
func (s *Store) GetFoodItems() []FoodItem {
	items := make([]FoodItem, 0)
	rows, err := s.db.Query("select id, manufacturer, product, protein, carbohydrate, fat, kcal , owner_id from food_items")
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
			&item.OwnerID,
		)
		items = append(items, item)
	}
	if err != nil {
		panic(err)
	}
	return items
}
