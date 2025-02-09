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
	err := s.db.QueryRow("select id, manufacturer, product from food_items").Scan(&item.ID, &item.Manufacturer)
	if err != nil {
		panic(err)
	}
	return &item
}
func (s *Store) GetFoodItems() []FoodItem {
	items := make([]FoodItem, 0)
	rows, err := s.db.Query("select fi.id, fi.manufacturer,fi.product, m.protein, m.carbohydrate, m.fat, m.kcal from food_items as fi join food_items_macronutrients as junction on fi.id = junction.food_item_id join macronutrients m on junction.macronutrient_id = m.id")
	for rows.Next() {
		item := FoodItem{}
		rows.Scan(&item.ID, &item.Manufacturer, &item.Product, &item.Macronutrients.Protein, &item.Macronutrients.Carbohydrate, &item.Macronutrients.Fat, &item.Macronutrients.KCal)
		items = append(items, item)
	}
	if err != nil {
		panic(err)
	}
	return items
}
