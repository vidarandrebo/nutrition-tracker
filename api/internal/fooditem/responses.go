package fooditem

type GetFoodItemResponse struct {
	ID           int64  `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
}
