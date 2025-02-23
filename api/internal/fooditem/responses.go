package fooditem

type FoodItemResponse struct {
	ID           int64   `json:"id"`
	Manufacturer string  `json:"manufacturer"`
	Product      string  `json:"product"`
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
	Fat          float64 `json:"fat"`
	KCal         float64 `json:"kCal"`
	Source       string  `json:"source,omitempty"`
}
