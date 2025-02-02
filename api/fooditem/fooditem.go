package fooditem

type FoodItem struct {
	Name string
	ID   int
}

func New(name string, id int) *FoodItem {
	return &FoodItem{
		Name: "",
		ID:   0,
	}
}
