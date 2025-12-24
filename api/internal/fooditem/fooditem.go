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
	Micronutrients []*Micronutrient
	PortionSizes   []*PortionSize
	Source         string
	OwnerID        int64
}

func NewFoodItem() *FoodItem {
	return &FoodItem{}
}

func (fi *FoodItem) ToResponse() api.FoodItemResponse {
	micronutrients := make([]api.FoodItemMicronutrientResponse, 0, len(fi.Micronutrients))
	for _, item := range fi.Micronutrients {
		micronutrients = append(micronutrients, item.ToResponse())
	}
	portionSizes := make([]api.FoodItemPortionSizeResponse, 0, len(fi.PortionSizes))
	for _, item := range fi.PortionSizes {
		portionSizes = append(portionSizes, item.ToResponse())
	}
	return api.FoodItemResponse{
		Carbohydrate:   fi.Carbohydrate,
		Fat:            fi.Fat,
		Id:             fi.ID,
		KCal:           fi.KCal,
		Manufacturer:   fi.Manufacturer,
		Product:        fi.Product,
		Protein:        fi.Protein,
		IsPublic:       fi.Public,
		Source:         fi.Source,
		OwnerId:        fi.OwnerID,
		Micronutrients: &micronutrients,
		PortionSizes:   &portionSizes,
	}
}

func (fi *FoodItem) FromRequest(r *api.FoodItemPostRequest) *FoodItem {
	kCal := 0.0
	if r.KCal == nil {
		kCal = r.Protein*4 + r.Carbohydrate*4 + r.Fat*9
	} else {
		kCal = *r.KCal
	}
	fi.Manufacturer = r.Manufacturer
	fi.Product = r.Product
	fi.Protein = r.Protein
	fi.Carbohydrate = r.Carbohydrate
	fi.Fat = r.Fat
	fi.KCal = kCal
	fi.Public = r.IsPublic
	fi.Micronutrients = make([]*Micronutrient, 0)
	fi.PortionSizes = make([]*PortionSize, 0)

	return fi
}

func (fi *FoodItem) HasAccess(userId int64) bool {
	if fi.Public || fi.OwnerID == userId {
		return true
	}
	return false
}

func (fi *FoodItem) FromMatvareTabellen(item matvaretabellen.Food) *FoodItem {
	macroNames := []string{"Protein", "Karbo", "Fett"}
	micronutrients := make([]*Micronutrient, 0)
	for _, constituent := range item.Constituents {
		if constituent.Quantity == 0.0 {
			continue
		}
		if !slices.Contains(macroNames, constituent.NutrientID) {
			micronutrients = append(micronutrients, &Micronutrient{
				Name:   constituent.NutrientID,
				Amount: CalcAmount(constituent.Quantity, constituent.Unit),
			})
		}
	}
	fi.ID = 0
	fi.Manufacturer = ""
	fi.Product = item.FoodName
	fi.Protein = item.Protein()
	fi.Carbohydrate = item.Carbohydrate()
	fi.Fat = item.Fat()
	fi.KCal = float64(item.Calories.Quantity)
	fi.Public = true
	fi.Micronutrients = micronutrients
	fi.Source = "matvaretabellen.no"
	fi.OwnerID = 0
	return fi
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

func (fi *FoodItem) ToTable() TableFoodItem {
	return TableFoodItem{
		ID:           fi.ID,
		Manufacturer: fi.Manufacturer,
		Product:      fi.Product,
		Protein:      fi.Protein,
		Carbohydrate: fi.Carbohydrate,
		Fat:          fi.Fat,
		KCal:         fi.KCal,
		Public:       fi.Public,
		Source:       fi.Source,
		OwnerID:      fi.OwnerID,
	}
}

func (fi *FoodItem) FromTable(item TableFoodItem) *FoodItem {
	fi.ID = item.ID
	fi.Manufacturer = item.Manufacturer
	fi.Product = item.Product
	fi.Protein = item.Protein
	fi.Carbohydrate = item.Carbohydrate
	fi.Fat = item.Fat
	fi.KCal = item.KCal
	fi.Public = item.Public
	fi.Micronutrients = make([]*Micronutrient, 0)
	fi.PortionSizes = make([]*PortionSize, 0)
	fi.Source = item.Source
	fi.OwnerID = item.OwnerID
	return fi
}
