package fooditem

import (
	"database/sql"
	"errors"
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

func fromFoodItemComplete(rows []TableFoodItemComplete) []FoodItem {
	portionSizes := make(map[int64]map[int64]PortionSize)
	micronutrients := make(map[int64]map[int64]Micronutrient)
	foodItems := make(map[int64]FoodItem)
	for _, item := range rows {
		_, ok := foodItems[item.FI.ID]
		if !ok {
			foodItems[item.FI.ID] = item.FI.ToFoodItem()
			portionSizes[item.FI.ID] = make(map[int64]PortionSize)
			micronutrients[item.FI.ID] = make(map[int64]Micronutrient)
		}
		ps, err := item.PS.ToPortionSize()
		if err == nil {
			if _, ok = portionSizes[item.FI.ID][ps.ID]; !ok {
				portionSizes[item.FI.ID][ps.ID] = ps
			}
		}
		_, ok = micronutrients[item.FI.ID][item.M.ID]
		if !ok {
			micronutrients[item.FI.ID][item.M.ID] = item.M.ToMicronutrient()
		}
	}
	out := make([]FoodItem, 0, len(foodItems))
	for _, item := range foodItems {
		for _, micronutrient := range micronutrients[item.ID] {
			item.Micronutrients = append(item.Micronutrients, micronutrient)
		}
		for _, portion := range portionSizes[item.ID] {
			item.PortionSizes = append(item.PortionSizes, portion)
		}
		out = append(out, item)
	}
	return out
}

func fromFoodItemAndPortion(rows []TableFoodItemAndPortion) []FoodItem {
	portionSizes := make(map[int64]map[int64]PortionSize)
	foodItems := make(map[int64]FoodItem)
	for _, item := range rows {
		_, ok := foodItems[item.FI.ID]
		if !ok {
			foodItems[item.FI.ID] = item.FI.ToFoodItem()
			portionSizes[item.FI.ID] = make(map[int64]PortionSize)
		}
		ps, err := item.P.ToPortionSize()
		if err == nil {
			if _, ok = portionSizes[item.FI.ID][ps.ID]; !ok {
				portionSizes[item.FI.ID][ps.ID] = ps
			}
		}
	}
	out := make([]FoodItem, 0, len(foodItems))
	for _, item := range foodItems {
		for _, portion := range portionSizes[item.ID] {
			item.PortionSizes = append(item.PortionSizes, portion)
		}
		out = append(out, item)
	}
	return out
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

func (tf TableFoodItem) ToFoodItem() FoodItem {
	return FoodItem{
		ID:             tf.ID,
		Manufacturer:   tf.Manufacturer,
		Product:        tf.Product,
		Protein:        tf.Protein,
		Carbohydrate:   tf.Carbohydrate,
		Fat:            tf.Fat,
		KCal:           tf.KCal,
		Public:         tf.Public,
		Micronutrients: make([]Micronutrient, 0),
		PortionSizes:   make([]PortionSize, 0),
		Source:         tf.Source,
		OwnerID:        tf.OwnerID,
	}
}

type TablePortionSize struct {
	ID     sql.NullInt64
	Amount float64
	Name   string
}

func (tp TablePortionSize) ToPortionSize() (PortionSize, error) {
	if tp.ID.Valid {
		return PortionSize{
			ID:     tp.ID.Int64,
			Name:   tp.Name,
			Amount: tp.Amount,
		}, nil
	}
	return PortionSize{}, errors.New("portion is null")
}

type TableMicronutrient struct {
	ID     int64
	Amount float64
	Name   string
}

func (tm TableMicronutrient) ToMicronutrient() Micronutrient {
	return Micronutrient{
		ID:     tm.ID,
		Name:   tm.Name,
		Amount: tm.Amount,
	}
}
