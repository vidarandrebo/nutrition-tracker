package fooditem

type TableFoodItem struct {
	ID           int64 `db: "db"`
	Manufacturer string
	Product      string
	Protein      float64
	Carbohydrate float64
	Fat          float64
	KCal         float64
	Public       bool
	Source       string
	OwnerID      int64
}

type TablePortionSize struct {
	ID         int64
	Amount     float64
	Name       string
	FoodItemID int64
}

type TableMicronutrient struct {
	ID         int64
	Amount     float64
	Name       string
	FoodItemID int64
}
