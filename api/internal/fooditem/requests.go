package fooditem

type PostFoodItemRequest struct {
	Manufacturer string  `json:"manufacturer"`
	Product      string  `json:"product"`
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
	Fat          float64 `json:"fat"`
	KCal         float64 `json:"kCal"`
	Public       bool    `json:"public"`
}

func (fr *PostFoodItemRequest) ToFoodItem() FoodItem {
	kCal := fr.KCal
	if fr.KCal == 0.0 {
		kCal = 4*fr.Protein + 4*fr.Carbohydrate + 9*fr.Fat
	}
	item := FoodItem{
		ID:             0,
		Manufacturer:   fr.Manufacturer,
		Product:        fr.Product,
		Protein:        fr.Protein,
		Carbohydrate:   fr.Carbohydrate,
		Fat:            fr.Fat,
		KCal:           kCal,
		Public:         false,
		Micronutrients: nil,
	}
	return item
}
