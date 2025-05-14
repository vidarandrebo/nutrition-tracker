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
	err = tx.QueryRow("insert into meals as m (sequence_number, meal_time, owner_id) values ($1, $2, $3) returning m.id", meal.SequenceNumber, meal.Timestamp, meal.OwnerID).Scan(&meal.ID)

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

func (s *Store) GetByDate(ownerID int64, dateFrom time.Time, dateTo time.Time) []Meal {
	rows, err := s.DB.Query("select m.id, m.meal_time, m.sequence_number, m.owner_id from meals m where m.owner_id = $1 and m.meal_time >= $2 and m.meal_time < $3", ownerID, dateFrom, dateTo)
	if err != nil {
		panic(err)
	}
	meals := make([]Meal, 0)
	for rows.Next() {
		meal := Meal{}
		rows.Scan(&meal.ID, &meal.Timestamp, &meal.SequenceNumber, &meal.OwnerID)
		meals = append(meals, meal)
	}
	return meals
}
func (s *Store) GetById(ownerID int64, id int64) (Meal, error) {
	row := s.DB.QueryRow("select m.id, m.meal_time, m.sequence_number, m.owner_id from meals m where m.id = $1 and m.owner_id = $2", id, ownerID)
	meal := Meal{}
	err := row.Scan(&meal.ID, &meal.Timestamp, &meal.SequenceNumber, &meal.OwnerID)

	if err != nil {
		return Meal{}, fmt.Errorf("no meal with id %d", id)
	}
	return meal, nil
}
