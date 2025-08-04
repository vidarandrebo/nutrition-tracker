package fooditem

import (
	"maps"
	"slices"
)

type TableFoodItemComplete struct {
	FI TableFoodItem
	M  TableMicronutrient
	PS TablePortionSize
}
type TableFoodItemAndPortion struct {
	FI TableFoodItem
	P  TablePortionSize
}

func fromFoodItemComplete(complete ...TableFoodItemComplete) []FoodItem {
	portionSizes := make(map[int64]map[int64]PortionSize)
	micronutrients := make(map[int64]map[int64]Micronutrient)
	foodItems := make(map[int64]FoodItem)
	for _, item := range complete {
		_, ok := foodItems[item.FI.ID]
		if !ok {
			foodItems[item.FI.ID] = FoodItem{
				ID:             item.FI.ID,
				Manufacturer:   item.FI.Manufacturer,
				Product:        item.FI.Product,
				Protein:        item.FI.Protein,
				Carbohydrate:   item.FI.Carbohydrate,
				Fat:            item.FI.Fat,
				KCal:           item.FI.KCal,
				Public:         item.FI.Public,
				Micronutrients: make([]Micronutrient, 0),
				PortionSizes:   make([]PortionSize, 0),
				Source:         item.FI.Source,
				OwnerID:        item.FI.OwnerID,
			}
			portionSizes[item.FI.ID] = make(map[int64]PortionSize)
			micronutrients[item.FI.ID] = make(map[int64]Micronutrient)
		}
		_, ok = portionSizes[item.FI.ID][item.PS.ID]
		if !ok {
			portionSizes[item.PS.ID][item.PS.ID] = PortionSize{
				ID:     item.PS.ID,
				Name:   item.PS.Name,
				Amount: item.PS.Amount,
			}
		}
		_, ok = micronutrients[item.FI.ID][item.M.ID]
		if !ok {
			portionSizes[item.FI.ID][item.M.ID] = PortionSize{
				ID:     item.M.ID,
				Name:   item.M.Name,
				Amount: item.M.Amount,
			}
		}
	}
	for _, item := range foodItems {
		for _, micronutrient := range micronutrients[item.ID] {
			item.Micronutrients = append(item.Micronutrients, Micronutrient{
				ID:     micronutrient.ID,
				Name:   micronutrient.Name,
				Amount: micronutrient.Amount,
			})
		}
		for _, portion := range portionSizes[item.ID] {
			item.PortionSizes = append(item.PortionSizes, PortionSize{
				ID:     portion.ID,
				Name:   portion.Name,
				Amount: portion.Amount,
			})
		}
	}
	return slices.Collect(maps.Values(foodItems))
}

type TableFoodItem struct {
	ID           int64
	Manufacturer string
	Product      string
	Protein      float64
	Carbohydrate float64
	Fat          float64
	KCal         float64
	Public       bool
	Source       string
	OwnerID      int64
}

type TablePortionSize struct {
	ID     int64
	Amount float64
	Name   string
}
type TableMicronutrient struct {
	ID     int64
	Amount float64
	Name   string
}
