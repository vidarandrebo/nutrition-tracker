package meal

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"
)

type Repository struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	return &Repository{
		DB:     db,
		Logger: logger,
	}
}

func (r *Repository) Add(meal Meal) (Meal, error) {
	tx, err := r.DB.Begin()
	err = tx.QueryRow(`
		INSERT INTO meals AS m (sequence_number, meal_time, owner_id) 
		VALUES ($1, $2, $3) 
		RETURNING m.id`,
		meal.SequenceNumber, meal.Timestamp, meal.OwnerID).Scan(&meal.ID)
	if err != nil {
		return Meal{}, err
	}

	for _, entry := range meal.Entries {
		err = tx.QueryRow(`
			INSERT INTO meal_entries AS me (amount, food_item_id, recipe_id, meal_id) 
			VALUES ($1, $2, $3) 
			RETURNING me.id`,
			entry.Amount,
			entry.foodItemID,
			entry.recipeID,
			meal.ID,
		).Scan(&entry.ID)
		if err != nil {
			return Meal{}, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return Meal{}, err
	}
	return meal, nil
}

func (r *Repository) GetByDate(ownerID int64, dateFrom time.Time, dateTo time.Time) ([]Meal, error) {
	rows, err := r.DB.Query(`
	WITH meal_for_day AS (
		SELECT id, meal_time, sequence_number, owner_id 
		FROM meals 
		WHERE owner_id = $1 
		  AND meal_time >= $2 
		  AND meal_time < $3
	) 
	SELECT m.id, m.meal_time, m.sequence_number, m.owner_id, me.id, me.food_item_id, me.recipe_id, me.amount
	FROM meal_for_day m 
		LEFT JOIN meal_entries me ON me.meal_id = m.id`,
		ownerID,
		dateFrom,
		dateTo,
	)
	if err != nil {
		return nil, err
	}
	meals := make([]Meal, 0)
	entries := make(map[int64][]Entry)
	lastMealId := int64(0)
	for rows.Next() {
		meal := Meal{}
		entry := Entry{}
		rows.Scan(&meal.ID, &meal.Timestamp, &meal.SequenceNumber, &meal.OwnerID, &entry.ID, &entry.foodItemID, &entry.recipeID, &entry.Amount)
		if lastMealId != meal.ID {
			meals = append(meals, meal)
			entries[meal.ID] = make([]Entry, 0)
		}
		if entry.IsValid() {
			entries[meal.ID] = append(entries[meal.ID], entry)
		} else {
			r.Logger.Error("failed to load entry", slog.Any("e", entry))
		}
		lastMealId = meal.ID
	}

	for i := 0; i < len(meals); i++ {
		meals[i].Entries = entries[meals[i].ID]
	}
	return meals, nil
}

func (r *Repository) GetById(id int64, ownerID int64) (Meal, error) {
	row := r.DB.QueryRow(`
		SELECT m.id, m.meal_time, m.sequence_number, m.owner_id 
		FROM meals m 
		WHERE m.id = $1 
		  AND m.owner_id = $2`,
		id, ownerID)
	meal := Meal{}
	err := row.Scan(&meal.ID, &meal.Timestamp, &meal.SequenceNumber, &meal.OwnerID)
	if err != nil {
		return Meal{}, fmt.Errorf("no meal with id %d", id)
	}
	return meal, nil
}

func (r *Repository) AddMealEntry(entry Entry, mealID int64, ownerID int64) (Entry, error) {
	// only an ownership check
	_, err := r.GetById(mealID, ownerID)
	if err != nil {
		return Entry{}, err
	}

	err = r.DB.QueryRow(`
		INSERT INTO meal_entries AS me (meal_id, food_item_id, recipe_id,amount) 
		VALUES ($1, $2, $3, $4) 
		RETURNING me.id`,
		mealID, entry.FoodItemID(), entry.RecipeID(), entry.Amount).Scan(&entry.ID)
	if err != nil {
		return Entry{}, err
	}

	return entry, nil
}

func (r *Repository) DeleteMeal(id int64, ownerID int64) error {
	_, err := r.DB.Query(`
		DELETE FROM meals WHERE id = $1 AND owner_id = $2
	`, id, ownerID)
	if err != nil {
		r.Logger.Error("failed to delete meal", slog.Int64("mealID", id), slog.Any("err", err))
		return err
	}
	return nil
}

func (r *Repository) DeleteMealEntry(entryID int64, mealID int64, ownerID int64) error {
	_, err := r.DB.Query(`
		WITH owned_meals AS (
			SELECT id
			FROM meals
			WHERE id = $2 
  			  AND owner_id = $3
		)
		DELETE FROM meal_entries
		WHERE id = $1 
		  AND meal_id IN (SELECT id FROM owned_meals)
	`, entryID, mealID, ownerID)
	if err != nil {
		r.Logger.Error("failed to delete meal entry", slog.Int64("mealID", entryID), slog.Any("err", err))
		return err
	}
	return nil
}
