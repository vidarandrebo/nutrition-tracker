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
		insert into meals as m (sequence_number, meal_time, owner_id) 
		values ($1, $2, $3) 
		returning m.id`,
		meal.SequenceNumber, meal.Timestamp, meal.OwnerID).Scan(&meal.ID)

	if err != nil {
		panic(err)
	}

	for _, entry := range meal.Entries {
		err = tx.QueryRow(`
			insert into meal_entries as me (amount, food_item_id, meal_id) 
			values ($1, $2, $3) 
			returning me.id`,
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
	rows, err := s.DB.Query(`
	with meal_for_day as (
		select id, meal_time, sequence_number, owner_id 
		from meals 
		where owner_id = $1 
		  and meal_time >= $2 
		  and meal_time < $3
	) 
	select m.id, m.meal_time, m.sequence_number, m.owner_id, me.id, me.food_item_id, me.amount
	from meal_for_day m 
		left join meal_entries me on me.meal_id = m.id`,
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
		rows.Scan(&meal.ID, &meal.Timestamp, &meal.SequenceNumber, &meal.OwnerID, &entry.ID, &entry.FoodItemID, &entry.Amount)
		if lastMealId != meal.ID {
			meals = append(meals, meal)
			entries[meal.ID] = make([]Entry, 0)
		}
		entries[meal.ID] = append(entries[meal.ID], entry)
		lastMealId = meal.ID
	}

	for i := 0; i < len(meals); i++ {
		meals[i].Entries = entries[meals[i].ID]
	}
	return meals
}
func (s *Store) GetById(id int64, ownerID int64) (Meal, error) {
	row := s.DB.QueryRow(`
		select m.id, m.meal_time, m.sequence_number, m.owner_id 
		from meals m 
		where m.id = $1 
		  and m.owner_id = $2`,
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
		insert into meal_entries as me (meal_id, food_item_id, amount) 
		values ($1, $2, $3) 
		returning me.id`,
		mealID, entry.FoodItemID, entry.Amount).Scan(&entry.ID)

	if err != nil {
		panic(err)
	}

	return entry, nil
}
