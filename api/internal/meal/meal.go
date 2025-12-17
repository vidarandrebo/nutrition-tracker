package meal

import (
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
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

func (m *Meal) ToResponse() api.MealResponse {
	foodItemEntries := make([]api.FoodItemMealEntryResponse, 0, len(m.FoodItemEntries))
	for _, fie := range m.FoodItemEntries {
		foodItemEntries = append(foodItemEntries, fie.ToResponse())
	}
	recipeEntries := make([]api.RecipeMealEntryResponse, 0, len(m.RecipeEntries))
	for _, re := range m.RecipeEntries {
		recipeEntries = append(recipeEntries, re.ToResponse())
	}
	macroEntries := make([]api.MacronutrientMealEntryResponse, 0, len(m.MacronutrientEntries))
	for _, me := range m.MacronutrientEntries {
		macroEntries = append(macroEntries, me.ToResponse())
	}
	return api.MealResponse{
		Id:                   m.ID,
		SequenceNumber:       m.SequenceNumber,
		Timestamp:            m.Timestamp,
		FoodItemEntries:      &foodItemEntries,
		MacronutrientEntries: &macroEntries,
		RecipeEntries:        &recipeEntries,
	}
}

func FromRequest(r api.PostMealRequest) *Meal {
	return &Meal{
		ID:                   0,
		SequenceNumber:       0,
		Timestamp:            r.Timestamp,
		FoodItemEntries:      nil,
		RecipeEntries:        nil,
		MacronutrientEntries: nil,
		OwnerID:              0,
	}
}

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

func RMEFromRequest(r api.PostRecipeMealEntryRequest) *RecipeMealEntry {
	return &RecipeMealEntry{
		ID:             0,
		RecipeID:       r.RecipeId,
		Amount:         r.Amount,
		SequenceNumber: 0,
	}
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

func (rme *RecipeMealEntry) ToResponse() api.RecipeMealEntryResponse {
	return api.RecipeMealEntryResponse{
		Id:             rme.ID,
		Amount:         rme.Amount,
		RecipeId:       rme.RecipeID,
		SequenceNumber: rme.SequenceNumber,
	}
}

type FoodItemMealEntry struct {
	ID             int64
	FoodItemID     int64
	Amount         float64
	SequenceNumber int
}

func FIMEFromRequest(r api.PostFoodItemMealEntryRequest) *FoodItemMealEntry {
	return &FoodItemMealEntry{
		ID:             0,
		FoodItemID:     r.FoodItemId,
		Amount:         r.Amount,
		SequenceNumber: 0,
	}
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

func (fime *FoodItemMealEntry) ToResponse() api.FoodItemMealEntryResponse {
	return api.FoodItemMealEntryResponse{
		Id:             fime.ID,
		Amount:         fime.Amount,
		FoodItemId:     fime.FoodItemID,
		SequenceNumber: fime.SequenceNumber,
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

func MNEFromRequest(r api.PostMacronutrientMealEntryRequest) *MacronutrientMealEntry {
	return &MacronutrientMealEntry{
		ID:             0,
		SequenceNumber: 0,
		Protein:        r.Protein,
		Carbohydrate:   r.Carbohydrate,
		Fat:            r.Fat,
		KCal:           r.Fat,
	}
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
func (mme *MacronutrientMealEntry) ToResponse() api.MacronutrientMealEntryResponse {
	return api.MacronutrientMealEntryResponse{
		Id:             mme.ID,
		Protein:        mme.Protein,
		Carbohydrate:   mme.Carbohydrate,
		Fat:            mme.Fat,
		KCal:           mme.KCal,
		SequenceNumber: mme.SequenceNumber,
	}

}
