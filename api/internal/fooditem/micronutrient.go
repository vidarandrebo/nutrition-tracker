package fooditem

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type Micronutrient struct {
	ID     int64
	Name   string
	Amount float64
}

func (mn *Micronutrient) ToResponse() api.MicronutrientResponse {
	return api.MicronutrientResponse{
		Amount: mn.Amount,
		Id:     mn.ID,
		Name:   mn.Name,
	}
}

func (mn *Micronutrient) ToTable(foodItemID int64) TableFoodItemMacronutrient {
	return TableFoodItemMacronutrient{
		ID:         mn.ID,
		Amount:     mn.Amount,
		Name:       mn.Name,
		FoodItemID: foodItemID,
	}
}

func FromMicronutrientTable(item TableFoodItemMacronutrient) *Micronutrient {
	return &Micronutrient{
		ID:     item.ID,
		Name:   item.Name,
		Amount: item.Amount,
	}
}

func FromMicronutrientPost(r *api.PostFoodItemMicronutrient) *Micronutrient {
	return &Micronutrient{
		Name:   r.Name,
		Amount: r.Amount,
	}
}
