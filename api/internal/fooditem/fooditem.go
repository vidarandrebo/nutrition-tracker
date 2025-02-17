package fooditem

type FoodItem struct {
	ID             int64
	Manufacturer   string
	Product        string
	Protein        float64
	Carbohydrate   float64
	Fat            float64
	KCal           float64
	Micronutrients []Micronutrient
	OwnerID        int64
}

func (fi FoodItem) ToFoodItemResponse() FoodItemResponse {
	return FoodItemResponse{
		ID:           fi.ID,
		Manufacturer: fi.Manufacturer,
		Product:      fi.Product,
		Protein:      fi.Protein,
		Carbohydrate: fi.Carbohydrate,
		Fat:          fi.Fat,
		KCal:         fi.KCal,
	}
}

type Micronutrient struct {
	ID     int64
	Name   string
	Amount float64
}
