package meal

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"
	"time"
)

type IRepository interface {
	Add(item TableMeal) (TableMeal, error)
	AddRecipeEntry(item TableMealRecipeEntry) (TableMealRecipeEntry, error)
	AddFoodItemEntry(item TableMealFoodItemEntry) (TableMealFoodItemEntry, error)
	AddMacronutrientEntry(item TableMealMacronutrientEntry) (TableMealMacronutrientEntry, error)
	Delete(id int64) error
	DeleteRecipeEntry(id int64, mealID int64) error
	DeleteFoodItemEntry(id int64, mealID int64) error
	DeleteMacronutrientEntry(id int64, mealID int64) error
	GetByDate(from time.Time, to time.Time, ownerID int64) ([]TableMeal, error)
	GetById(id int64) (TableMeal, error)
	GetFoodItemEntries(mealID int64) ([]TableMealFoodItemEntry, error)
	GetMacronutrientEntries(mealID int64) ([]TableMealMacronutrientEntry, error)
	GetRecipeEntries(mealID int64) ([]TableMealRecipeEntry, error)
}

type Repository struct {
	db     *sql.DB
	logger *slog.Logger
}

// DeleteRecipeEntry returns the database error if deleting the recipe entry fails.
// It returns no error when no entry matches the ID and meal ID.
func (r *Repository) DeleteRecipeEntry(id int64, mealID int64) error {
	_, err := r.db.Exec(`
		DELETE FROM meal_recipe_entries re
		WHERE re.id = $1 AND re.meal_id = $2
		`,
		id,
		mealID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFoodItemEntry returns the database error if deleting the food item entry fails.
// It returns no error when no entry matches the ID and meal ID.
func (r *Repository) DeleteFoodItemEntry(id int64, mealID int64) error {
	_, err := r.db.Exec(`
		DELETE FROM meal_food_item_entries fie
		WHERE fie.id = $1 AND fie.meal_id = $2
		`,
		id,
		mealID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMacronutrientEntry returns the database error if deleting the macronutrient entry fails.
// It returns no error when no entry matches the ID and meal ID.
func (r *Repository) DeleteMacronutrientEntry(id int64, mealID int64) error {
	_, err := r.db.Exec(`
		DELETE FROM meal_macronutrient_entries me
		WHERE me.id = $1 AND me.meal_id = $2
		`,
		id,
		mealID,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	r := Repository{db: db}
	r.logger = logger.With(slog.Any("module", reflect.TypeOf(r)))
	return &r
}

// Add returns the database or scan error if inserting the meal fails,
// including sql.ErrNoRows when the insert returns no row.
func (r *Repository) Add(item TableMeal) (TableMeal, error) {
	err := r.db.QueryRow(`
		INSERT INTO meals AS m (meal_time, sequence_number, owner_id)
		VALUES  ($1,$2,$3)
		RETURNING m.id`,
		item.MealTime,
		item.SequenceNumber,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		return TableMeal{}, err
	}
	r.logger.Info("added new meal", slog.Any("meal", item))

	return item, nil
}

// AddRecipeEntry returns the database or scan error if inserting the recipe
// entry fails, including sql.ErrNoRows when the insert returns no row.
func (r *Repository) AddRecipeEntry(item TableMealRecipeEntry) (TableMealRecipeEntry, error) {
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
		return TableMealRecipeEntry{}, err
	}
	r.logger.Info("added recipe entry to meal", slog.Any("entry", item))

	return item, nil
}

// AddFoodItemEntry returns the database or scan error if inserting the food
// item entry fails, including sql.ErrNoRows when the insert returns no row.
func (r *Repository) AddFoodItemEntry(item TableMealFoodItemEntry) (TableMealFoodItemEntry, error) {
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
		return TableMealFoodItemEntry{}, err
	}
	r.logger.Info("added food item entry to meal", slog.Any("entry", item))

	return item, nil
}

// AddMacronutrientEntry returns the database or scan error if inserting the
// macronutrient entry fails, including sql.ErrNoRows when the insert returns no row.
func (r *Repository) AddMacronutrientEntry(item TableMealMacronutrientEntry) (TableMealMacronutrientEntry, error) {
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
		return TableMealMacronutrientEntry{}, err
	}
	r.logger.Info("added macronutrient entry to meal", slog.Any("entry", item))

	return item, nil
}

// Delete returns the database error if deleting the meal fails.
// It returns no error when the meal does not exist.
func (r *Repository) Delete(id int64) error {
	_, err := r.db.Exec(`
		DELETE FROM meals 
		WHERE id = $1
	`, id,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetByDate returns database, scan, iteration, or close errors without partial
// results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no meals match.
func (r *Repository) GetByDate(from time.Time, to time.Time, ownerID int64) (items []TableMeal, err error) {
	items = make([]TableMeal, 0)
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
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			items, err = nil, errors.Join(err, closeErr)
		}
	}()
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
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// GetById returns sql.ErrNoRows if no meal matches, or the database or scan
// error if querying or scanning the meal fails.
func (r *Repository) GetById(id int64) (TableMeal, error) {
	item := TableMeal{}
	err := r.db.QueryRow(`
		SELECT m.id, m.sequence_number, m.meal_time, m.date_created, m.date_modified, m.owner_id
		FROM meals m
		WHERE m.id = $1
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
		return TableMeal{}, err
	}
	return item, nil
}

// GetFoodItemEntries returns database, scan, iteration, or close errors without
// partial results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no food item entries match.
func (r *Repository) GetFoodItemEntries(mealID int64) (items []TableMealFoodItemEntry, err error) {
	items = make([]TableMealFoodItemEntry, 0)
	rows, err := r.db.Query(`
		SELECT fime.id, fime.food_item_id, fime.amount, fime.sequence_number, fime.date_created, fime.date_modified, fime.meal_id
		FROM meal_food_item_entries fime
		WHERE fime.meal_id = $1
		`,
		mealID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			items, err = nil, errors.Join(err, closeErr)
		}
	}()

	for rows.Next() {
		item := TableMealFoodItemEntry{}
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
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// GetMacronutrientEntries returns database, scan, iteration, or close errors
// without partial results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no macronutrient entries match.
func (r *Repository) GetMacronutrientEntries(mealID int64) (items []TableMealMacronutrientEntry, err error) {
	items = make([]TableMealMacronutrientEntry, 0)
	rows, err := r.db.Query(`
		SELECT mme.id, mme.sequence_number,mme.protein, mme.carbohydrate, mme.fat, mme.kcal, mme.date_created, mme.date_modified, mme.meal_id
		FROM meal_macronutrient_entries mme
		WHERE mme.meal_id = $1
		`,
		mealID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			items, err = nil, errors.Join(err, closeErr)
		}
	}()

	for rows.Next() {
		item := TableMealMacronutrientEntry{}
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
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// GetRecipeEntries returns database, scan, iteration, or close errors without
// partial results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no recipe entries match.
func (r *Repository) GetRecipeEntries(mealID int64) (items []TableMealRecipeEntry, err error) {
	items = make([]TableMealRecipeEntry, 0)
	rows, err := r.db.Query(`
		SELECT rme.id, rme.recipe_id, rme.amount, rme.sequence_number, rme.date_created, rme.date_modified, rme.meal_id
		FROM meal_recipe_entries rme
		WHERE rme.meal_id = $1
		`,
		mealID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			items, err = nil, errors.Join(err, closeErr)
		}
	}()

	for rows.Next() {
		item := TableMealRecipeEntry{}
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
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
