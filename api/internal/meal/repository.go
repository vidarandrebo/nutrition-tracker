package meal

import (
	"database/sql"
	"log/slog"
	"reflect"
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IRepository interface {
	Add(item TableMeal) (TableMeal, error)
	AddRecipeEntry(item TableRecipeMealEntry) (TableRecipeMealEntry, error)
	AddFoodItemEntry(item TableFoodItemMealEntry) (TableFoodItemMealEntry, error)
	AddMacronutrientEntry(item TableMacronutrientMealEntry) (TableMacronutrientMealEntry, error)
	Delete(id int64) error
	GetByDate(from time.Time, to time.Time, ownerID int64) ([]TableMeal, error)
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
		VALUES  ($1,$2,$3)
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
 		INSERT INTO meal_recipe_entries AS rme (recipe_id, amount, sequence_number, meal_id)
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
        INSERT INTO meal_food_item_entries AS fime (food_item_id, amount, sequence_number, meal_id)
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
		INSERT INTO meal_macronutrient_entries AS mme (protein, carbohydrate, fat, carbohydrate, sequence_number, meal_id)
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
	_, err := r.db.Query(`
		DELETE FROM meals 
		WHERE id = $1
	`, id,
	)
	if err != nil {
		r.logger.Error("failed to delete meal", slog.Int64("mealID", id))
		return err
	}

	return nil
}

func (r Repository) GetByDate(from time.Time, to time.Time, ownerID int64) ([]TableMeal, error) {
	items := make([]TableMeal, 0)
	rows, err := r.db.Query(`
		WITH owned_meals AS (
		    SELECT *
			FROM meals 
			WHERE owner_id = $1
		)	
		SELECT m.id, m.sequence_number, m.meal_time, m.date_created, m.date_modified, m.owner_id
		FROM owned_meals m
		WHERE m.meal_time > $2
		  AND m.meal_time < $3
		`,
		ownerID,
		from,
		to,
	)
	if err != nil {
		r.logger.Error("failed to query rows of meals", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TableMeal{}
		err = rows.Scan(
			&item.ID,
			&item.SequenceNumber,
			&item.MealTime,
			&item.DateCreated,
			&item.DateModified,
			&item.OwnerID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of meals", slog.Any("err", err))
		}
		items = append(items, item)
	}
	return items, nil
}

func (r Repository) GetById(id int64) (TableMeal, error) {
	item := TableMeal{}
	err := r.db.QueryRow(`
		SELECT m.id, m.sequence_number, m.meal_time, m.date_created, m.date_modified, m.owner_id
		FROM meals m
		WHERE m.owner_id = $1
		`,
		id,
	).Scan(
		&item.ID,
		&item.SequenceNumber,
		&item.MealTime,
		&item.DateCreated,
		&item.DateModified,
		&item.OwnerID,
	)
	if err != nil {
		r.logger.Error("failed to query rows of meals", slog.Any("err", err))
		return TableMeal{}, err
	}
	return item, nil
}

func (r Repository) GetFoodItemEntries(mealID int64) ([]TableFoodItemMealEntry, error) {
	items := make([]TableFoodItemMealEntry, 0)
	rows, err := r.db.Query(`
		SELECT fime.id, fime.food_item_id, fime.amount, fime.sequence_number, fime.date_created, fime.date_modified, fime.meal_id
		FROM meal_food_item_entries fime
		WHERE fime.meal_id = $1
		`,
		mealID)
	if err != nil {
		r.logger.Error("failed to query rows of meal entries", slog.Any("err", err))
		return nil, err
	}

	for rows.Next() {
		item := TableFoodItemMealEntry{}
		err = rows.Scan(
			&item.ID,
			&item.FoodItemID,
			&item.Amount,
			&item.SequenceNumber,
			&item.DateCreated,
			&item.DateModified,
			&item.MealID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of meal entries", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r Repository) GetMacronutrientEntries(mealID int64) ([]TableMacronutrientMealEntry, error) {
	items := make([]TableMacronutrientMealEntry, 0)
	rows, err := r.db.Query(`
		SELECT mme.id, mme.sequence_number,mme.protein, mme.carbohydrate, mme.fat, mme.kcal, mme.date_created, mme.date_modified, mme.meal_id
		FROM meal_macronutrient_entries mme
		WHERE mme.meal_id = $1
		`,
		mealID)
	if err != nil {
		r.logger.Error("failed to query rows of meal entries", slog.Any("err", err))
		return nil, err
	}

	for rows.Next() {
		item := TableMacronutrientMealEntry{}
		err = rows.Scan(
			&item.ID,
			&item.SequenceNumber,
			&item.Protein,
			&item.Carbohydrate,
			&item.Fat,
			&item.KCal,
			&item.DateCreated,
			&item.DateModified,
			&item.MealID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of meal entries", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r Repository) GetRecipeEntries(mealID int64) ([]TableRecipeMealEntry, error) {
	items := make([]TableRecipeMealEntry, 0)
	rows, err := r.db.Query(`
		SELECT rme.id, rme.recipe_id, rme.amount, rme.sequence_number, rme.date_created, rme.date_modified, rme.meal_id
		FROM meal_recipe_entries rme
		WHERE rme.meal_id = $1
		`,
		mealID)
	if err != nil {
		r.logger.Error("failed to query rows of meal entries", slog.Any("err", err))
		return nil, err
	}

	for rows.Next() {
		item := TableRecipeMealEntry{}
		err = rows.Scan(
			&item.ID,
			&item.RecipeID,
			&item.Amount,
			&item.SequenceNumber,
			&item.DateCreated,
			&item.DateModified,
			&item.MealID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of meal entries", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r Repository) CheckOwnership(id int64, ownerID int64) error {
	meal := TableMeal{}
	err := r.db.QueryRow(`
		SELECT id, owner_id 
		FROM meals 
		WHERE id = $1
	`, id).Scan(
		&meal.ID,
		&meal.OwnerID,
	)
	if err != nil {
		return utils.ErrEntityNotFound
	}

	if meal.OwnerID == ownerID {
		return nil
	}

	return utils.ErrEntityNotOwned
}
