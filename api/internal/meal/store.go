package meal

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"
)

type Store struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func NewStore(db *sql.DB, logger *slog.Logger) *Store {
	return &Store{
		DB:     db,
		Logger: logger,
	}
}

func (s *Store) Add(meal Meal) Meal {
	tx, err := s.DB.Begin()
	err = tx.QueryRow(`
		INSERT INTO meals AS m (sequence_number, meal_time, owner_id) 
		VALUES ($1, $2, $3) 
		RETURNING m.id`,
		meal.SequenceNumber, meal.Timestamp, meal.OwnerID).Scan(&meal.ID)

	if err != nil {
		panic(err)
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
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return meal
}

func (s *Store) GetByDate(ownerID int64, dateFrom time.Time, dateTo time.Time) []Meal {
	rows, err := s.DB.Query(`
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
		panic(err)
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
			s.Logger.Error("failed to load entry", slog.Any("e", entry))
		}
		lastMealId = meal.ID
	}

	for i := 0; i < len(meals); i++ {
		meals[i].Entries = entries[meals[i].ID]
	}
	return meals
}
func (s *Store) GetById(id int64, ownerID int64) (Meal, error) {
	row := s.DB.QueryRow(`
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
func (s *Store) AddMealEntry(entry Entry, mealID int64, ownerID int64) (Entry, error) {
	// only an ownership check
	_, err := s.GetById(mealID, ownerID)
	if err != nil {
		return Entry{}, err
	}

	err = s.DB.QueryRow(`
		INSERT INTO meal_entries AS me (meal_id, food_item_id, recipe_id,amount) 
		VALUES ($1, $2, $3, $4) 
		RETURNING me.id`,
		mealID, entry.FoodItemID(), entry.RecipeID(), entry.Amount).Scan(&entry.ID)

	if err != nil {
		panic(err)
	}

	return entry, nil
}
