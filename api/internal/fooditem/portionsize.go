package fooditem

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type PortionSize struct {
	ID     int64
	Name   string
	Amount float64
}

func (ps *PortionSize) ToResponse() api.PortionSizeResponse {
	return api.PortionSizeResponse{
		Amount: ps.Amount,
		Id:     ps.ID,
		Name:   ps.Name,
	}
}

func (ps *PortionSize) ToTable(foodItemID int64) TablePortionSize {
	return TablePortionSize{
		ID:         ps.ID,
		Amount:     ps.Amount,
		Name:       ps.Name,
		FoodItemID: foodItemID,
	}
}

func FromPortionSizeTable(item TablePortionSize) *PortionSize {
	return &PortionSize{
		ID:     item.ID,
		Name:   item.Name,
		Amount: item.Amount,
	}
}

func FromPortionSizePost(r *api.PostFoodItemPortion) *PortionSize {
	return &PortionSize{
		Name:   r.Name,
		Amount: r.Amount,
	}
}
