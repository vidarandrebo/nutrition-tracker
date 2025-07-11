package fooditem

import (
	"slices"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/matvaretabellen"
)

type FoodItem struct {
	ID             int64
	Manufacturer   string
	Product        string
	Protein        float64
	Carbohydrate   float64
	Fat            float64
	KCal           float64
	Public         bool
	Micronutrients []Micronutrient
	Source         string
	OwnerID        int64
}

func (fi FoodItem) ToFoodItemResponse() api.FoodItemResponse {
	return api.FoodItemResponse{
		Carbohydrate: fi.Carbohydrate,
		Fat:          fi.Fat,
		Id:           fi.ID,
		KCal:         fi.KCal,
		Manufacturer: fi.Manufacturer,
		Product:      fi.Product,
		Protein:      fi.Protein,
		IsPublic:     fi.Public,
		Source:       fi.Source,
	}
}

func (fi FoodItem) HasAccess(userId int64) bool {
	if fi.Public || fi.OwnerID == userId {
		return true
	}
	return false
}

func FromMatvareTabellen(item matvaretabellen.Food) FoodItem {
	macroNames := []string{"Protein", "Karbo", "Fett"}
	micronutrients := make([]Micronutrient, 0)
	for _, constituent := range item.Constituents {
		if constituent.Quantity == 0.0 {
			continue
		}
		if !slices.Contains(macroNames, constituent.NutrientID) {
			micronutrients = append(micronutrients, Micronutrient{
				Name:   constituent.NutrientID,
				Amount: CalcAmount(constituent.Quantity, constituent.Unit),
			})
		}
	}
	foodItem := FoodItem{
		ID:             0,
		Manufacturer:   "",
		Product:        item.FoodName,
		Protein:        item.Protein(),
		Carbohydrate:   item.Carbohydrate(),
		Fat:            item.Fat(),
		KCal:           float64(item.Calories.Quantity),
		Public:         true,
		Micronutrients: micronutrients,
		Source:         "matvaretabellen.no",
		OwnerID:        0,
	}
	return foodItem
}

func CalcAmount(amount float64, unit string) float64 {
	switch unit {
	case "mg":
		return amount / 1000
	case "mg-ATE":
		return amount / 1000
	case "g":
		return amount
	case "\u00b5g":
		// μg
		return amount / 1000000
	case "\u00b5g-RE":
		// μg-RE
		return amount / 1000000
	case "":
		return amount
	default:
		panic(unit)
	}
}

type Micronutrient struct {
	ID     int64
	Name   string
	Amount float64
}
