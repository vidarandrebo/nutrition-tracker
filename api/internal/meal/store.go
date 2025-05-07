package meal

import (
	"database/sql"
	"log/slog"
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
	err = tx.QueryRow("insert into meals as m (sequence_number, meal_time) values ($1, $2) returning m.id", meal.SequenceNumber, meal.Timestamp).Scan(&meal.ID)

	if err != nil {
		panic(err)
	}

	for _, entry := range meal.Entries {
		err = tx.QueryRow("insert into meal_entries as me (amount, food_item_id, meal_id) values ($1, $2, $3) returning me.id",
			entry.Amount,
			entry.FoodItemIDOrNil(),
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
