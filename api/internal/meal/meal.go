package meal

import (
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Meal struct {
	ID                   int64
	SequenceNumber       int
	Timestamp            time.Time
	FoodItemEntries      []*FoodItemEntry
	RecipeEntries        []*RecipeEntry
	MacronutrientEntries []*MacronutrientEntry
	OwnerID              int64
}

func NewMeal() *Meal {
	return &Meal{}
}

func (m *Meal) ToResponse() api.MealResponse {
	foodItemEntries := make([]api.MealFoodItemEntryResponse, 0, len(m.FoodItemEntries))
	for _, fie := range m.FoodItemEntries {
		foodItemEntries = append(foodItemEntries, fie.ToResponse())
	}
	recipeEntries := make([]api.MealRecipeEntryResponse, 0, len(m.RecipeEntries))
	for _, re := range m.RecipeEntries {
		recipeEntries = append(recipeEntries, re.ToResponse())
	}
	macroEntries := make([]api.MealMacronutrientEntryResponse, 0, len(m.MacronutrientEntries))
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

func FromRequest(r api.MealPostRequest) *Meal {
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
		OwnerID:        m.OwnerID,
	}
}

func (m *Meal) FromTable(meal TableMeal) *Meal {
	m.ID = meal.ID
	m.SequenceNumber = meal.SequenceNumber
	m.Timestamp = meal.MealTime
	m.FoodItemEntries = make([]*FoodItemEntry, 0)
	m.RecipeEntries = make([]*RecipeEntry, 0)
	m.MacronutrientEntries = make([]*MacronutrientEntry, 0)
	m.OwnerID = meal.OwnerID
	return m
}
