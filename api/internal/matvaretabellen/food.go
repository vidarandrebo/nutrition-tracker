package matvaretabellen

type Foods struct {
	Items  []Food `json:"foods"`
	Locale string `json:"locale"`
}
type Food struct {
	SearchKeywords []string      `json:"searchKeywords"`
	Calories       Calories      `json:"calories"`
	Portions       []Portion     `json:"portions"`
	EdiblePart     EdiblePart    `json:"ediblePart,omitempty"`
	LangualCodes   []string      `json:"langualCodes"`
	Energy         Energy        `json:"energy,omitempty"`
	FoodName       string        `json:"foodName"`
	LatinName      string        `json:"latinName"`
	Constituents   []Constituent `json:"constituents"`
	URI            string        `json:"uri"`
	FoodGroupID    string        `json:"foodGroupId"`
	FoodID         string        `json:"foodId"`
}

func (f Food) Protein() float64 {
	for _, item := range f.Constituents {
		if item.NutrientID == "Protein" {
			return item.Quantity
		}
	}
	return 0.0
}

func (f Food) Carbohydrate() float64 {
	for _, item := range f.Constituents {
		if item.NutrientID == "Karbo" {
			return item.Quantity
		}
	}
	return 0.0
}

func (f Food) Fat() float64 {
	for _, item := range f.Constituents {
		if item.NutrientID == "Fett" {
			return item.Quantity
		}
	}
	return 0.0
}

type Calories struct {
	SourceID string `json:"sourceId"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
type Portion struct {
	PortionName string  `json:"portionName"`
	PortionUnit string  `json:"portionUnit"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
}
type EdiblePart struct {
	Percent  int    `json:"percent"`
	SourceID string `json:"sourceId"`
}
type Energy struct {
	SourceID string  `json:"sourceId"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
type Constituent struct {
	SourceID   string  `json:"sourceId"`
	NutrientID string  `json:"nutrientId"`
	Quantity   float64 `json:"quantity,omitempty"`
	Unit       string  `json:"unit,omitempty"`
}
