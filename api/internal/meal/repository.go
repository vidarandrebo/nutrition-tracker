package meal

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type IRepository interface {
	Add(item TableMeal) (TableMeal, error)
	AddRecipeEntry(item TableRecipeMealEntry) (TableRecipeMealEntry, error)
	AddFoodItemEntry(item TableFoodItemMealEntry) (TableFoodItemMealEntry, error)
	AddMacronutrientEntry(item TableMacronutrientMealEntry) (TableMacronutrientMealEntry, error)
	Delete(id int64) error
	Get(ownerID int64) ([]TableMeal, error)
	GetById(id int64) (TableMeal, error)
	GetFoodItemEntries(mealID int64) ([]TableFoodItemMealEntry, error)
	GetMacronutrientEntries(mealID int64) ([]TableMacronutrientMealEntry, error)
	GetRecipeEntries(mealID int64) ([]TableRecipeMealEntry, error)
	CheckOwnership(id int64, ownerID int64) error
}

type Repository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	r := Repository{db: db}
	r.logger = logger.With(slog.Any("module", reflect.TypeOf(r)))
	return &r
}

func (r Repository) Add(item TableMeal) (TableMeal, error) {
	err := r.db.QueryRow(`
		INSERT INTO meals AS m (meal_time, sequence_number, owner_id)
		VALUES  ($1,$2,3)
		RETURNING m.id`,
		item.MealTime,
		item.SequenceNumber,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add meal", slog.Int64("userID", item.OwnerID), slog.Any("err", err))
		return TableMeal{}, err
	}
	r.logger.Info("added new meal", slog.Any("meal", item))

	return item, nil
}

func (r Repository) AddRecipeEntry(item TableRecipeMealEntry) (TableRecipeMealEntry, error) {
	err := r.db.QueryRow(`
 		INSERT INTO recipe_meal_entries AS rme (recipe_id, amount, sequence_number, meal_id)
		VALUES ($1,$2,$3,$4)
		RETURNING rme.id`,
		item.RecipeID,
		item.Amount,
		item.SequenceNumber,
		item.MealID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add recipe entry to meal", slog.Int64("mealID", item.MealID), slog.Any("err", err))
		return TableRecipeMealEntry{}, err
	}
	r.logger.Info("added recipe entry to meal", slog.Any("entry", item))

	return item, nil
}

func (r Repository) AddFoodItemEntry(item TableFoodItemMealEntry) (TableFoodItemMealEntry, error) {
	err := r.db.QueryRow(`
        INSERT INTO food_item_meal_entries AS fime (food_item_id, amount, sequence_number, meal_id)
		VALUES ($1, $2, $3, $4)
		RETURNING fime.id
		`,
		item.FoodItemID,
		item.Amount,
		item.SequenceNumber,
		item.MealID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add food item entry to meal", slog.Int64("mealID", item.MealID), slog.Any("err", err))
		return TableFoodItemMealEntry{}, err
	}
	r.logger.Info("added food item entry to meal", slog.Any("entry", item))

	return item, nil

}

func (r Repository) AddMacronutrientEntry(item TableMacronutrientMealEntry) (TableMacronutrientMealEntry, error) {
	err := r.db.QueryRow(`
		INSERT INTO macronutrient_meal_entries AS mme (protein, carbohydrate, fat, carbohydrate, sequence_number, meal_id)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING mme.id
		`,
		item.Protein,
		item.Carbohydrate,
		item.Fat,
		item.KCal,
		item.SequenceNumber,
		item.MealID,
	).Scan(&item.ID)

	if err != nil {
		r.logger.Error("failed to add macronutrient entry to meal", slog.Int64("mealID", item.MealID), slog.Any("err", err))
		return TableMacronutrientMealEntry{}, err
	}
	r.logger.Info("added macronutrient entry to meal", slog.Any("entry", item))

	return item, nil

}

func (r Repository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Get(ownerID int64) ([]TableMeal, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetById(id int64) (TableMeal, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetFoodItemEntries(mealID int64) ([]TableFoodItemMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetMacronutrientEntries(mealID int64) ([]TableMacronutrientMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetRecipeEntries(mealID int64) ([]TableRecipeMealEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CheckOwnership(id int64, ownerID int64) error {
	//TODO implement me
	panic("implement me")
}
