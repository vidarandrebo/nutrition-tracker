package fooditem

type PostFoodItemRequest struct {
	Manufacturer string  `json:"manufacturer"`
	Product      string  `json:"product"`
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
	Fat          float64 `json:"fat"`
	KCal         float64 `json:"kCal"`
}

func (fr *PostFoodItemRequest) ToFoodItem() FoodItem {
	item := FoodItem{
		ID:             0,
		Manufacturer:   fr.Manufacturer,
		Product:        fr.Product,
		Protein:        fr.Protein,
		Carbohydrate:   fr.Carbohydrate,
		Fat:            fr.Fat,
		KCal:           fr.KCal,
		Micronutrients: nil,
	}
	return item
}
