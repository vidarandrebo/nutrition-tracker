package fooditem

type FoodItem struct {
	ID             int64
	Manufacturer   string
	Product        string
	Macronutrients Macronutrients
}

type Macronutrients struct {
	Protein      float64
	Carbohydrate float64
	Fat          float64
	KCal         float64
}
