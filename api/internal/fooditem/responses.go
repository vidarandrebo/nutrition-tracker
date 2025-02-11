package fooditem

type GetFoodItemResponse struct {
	ID             int64                    `json:"id"`
	Manufacturer   string                   `json:"manufacturer"`
	Product        string                   `json:"product"`
	Macronutrients GetMacronutrientResponse `json:"macronutrients"`
}

type PostFoodItemResponse GetFoodItemResponse

type GetMacronutrientResponse struct {
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
	Fat          float64 `json:"fat"`
	KCal         float64 `json:"kCal"`
}
