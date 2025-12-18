package meal

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type MacronutrientEntry struct {
	ID             int64
	SequenceNumber int
	Protein        float64
	Carbohydrate   float64
	Fat            float64
	KCal           float64
}

func NewMacronutrientEntry() *MacronutrientEntry {
	return &MacronutrientEntry{}
}

func MNEFromRequest(r api.PostMacronutrientMealEntryRequest) *MacronutrientEntry {
	return &MacronutrientEntry{
		ID:             0,
		SequenceNumber: 0,
		Protein:        r.Protein,
		Carbohydrate:   r.Carbohydrate,
		Fat:            r.Fat,
		KCal:           r.Fat,
	}
}

func (mme *MacronutrientEntry) ToTable(mealID int64) TableMealMacronutrientEntry {
	return TableMealMacronutrientEntry{
		ID:             mme.ID,
		SequenceNumber: mme.SequenceNumber,
		Protein:        mme.Protein,
		Carbohydrate:   mme.Carbohydrate,
		Fat:            mme.Fat,
		KCal:           mme.KCal,
		MealID:         mealID,
	}
}

func (mme *MacronutrientEntry) ToResponse() api.MacronutrientMealEntryResponse {
	return api.MacronutrientMealEntryResponse{
		Id:             mme.ID,
		Protein:        mme.Protein,
		Carbohydrate:   mme.Carbohydrate,
		Fat:            mme.Fat,
		KCal:           mme.KCal,
		SequenceNumber: mme.SequenceNumber,
	}
}

func (mme *MacronutrientEntry) FromTable(entry TableMealMacronutrientEntry) *MacronutrientEntry {
	mme.ID = entry.ID
	mme.SequenceNumber = entry.SequenceNumber
	mme.Protein = entry.Protein
	mme.Carbohydrate = entry.Carbohydrate
	mme.Fat = entry.Fat
	mme.KCal = entry.KCal
	return mme
}
