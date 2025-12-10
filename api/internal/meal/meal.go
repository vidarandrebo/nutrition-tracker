package meal

import (
	"time"
)

type Meal struct {
	ID                   int64
	SequenceNumber       int
	Timestamp            time.Time
	FoodItemEntries      []*FoodItemMealEntry
	RecipeEntries        []*RecipeMealEntry
	MacronutrientEntries []*MacronutrientMealEntry
	OwnerID              int64
}

//func (m Meal) ToResponse() api.MealResponse {
//	entries := make([]api.MealEntryResponse, 0, len(m.Entries))
//
//	for _, e := range m.Entries {
//		entries = append(entries, e.ToResponse())
//	}
//	return api.MealResponse{
//		Id:             m.ID,
//		SequenceNumber: m.SequenceNumber,
//		Timestamp:      m.Timestamp,
//		Entries:        entries,
//	}
//}

//func FromRequest(r api.PostMealRequest) *Meal {
//	return &Meal{
//		Timestamp:       r.Timestamp,
//		FoodItemEntries: make([]Entry[*fooditem.FoodItem], 0),
//		RecipeEntries:   make([]Entry[*recipe.Recipe], 0),
//		OwnerID:         0,
//	}
//}

func (m *Meal) ToTable() TableMeal {
	return TableMeal{
		ID:             m.ID,
		SequenceNumber: m.SequenceNumber,
		MealTime:       m.Timestamp,
		OwnerID:        0,
	}
}

type RecipeMealEntry struct {
	ID             int64
	RecipeID       int64
	Amount         float64
	SequenceNumber int
}

func (rme *RecipeMealEntry) ToTable(mealID int64) TableRecipeMealEntry {
	return TableRecipeMealEntry{
		ID:             rme.ID,
		RecipeID:       rme.RecipeID,
		Amount:         rme.Amount,
		SequenceNumber: rme.SequenceNumber,
		MealID:         mealID,
	}
}

type FoodItemMealEntry struct {
	ID             int64
	FoodItemID     int64
	Amount         float64
	SequenceNumber int
}

func (fime *FoodItemMealEntry) ToTable(mealID int64) TableFoodItemMealEntry {
	return TableFoodItemMealEntry{
		ID:             fime.ID,
		FoodItemID:     fime.FoodItemID,
		Amount:         fime.Amount,
		SequenceNumber: fime.SequenceNumber,
		MealID:         mealID,
	}
}

type MacronutrientMealEntry struct {
	ID             int64
	SequenceNumber int
	Protein        float64
	Carbohydrate   float64
	Fat            float64
	KCal           float64
}

func (mme *MacronutrientMealEntry) ToTable(mealID int64) TableMacronutrientMealEntry {
	return TableMacronutrientMealEntry{
		ID:             mme.ID,
		SequenceNumber: mme.SequenceNumber,
		Protein:        mme.Protein,
		Carbohydrate:   mme.Carbohydrate,
		Fat:            mme.Fat,
		KCal:           mme.KCal,
		MealID:         mealID,
	}
}
