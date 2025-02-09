package fooditem

type PostFoodItemRequest struct {
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
}
