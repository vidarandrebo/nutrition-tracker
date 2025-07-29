package recipe

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type Store struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewStore(db *sql.DB, logger *slog.Logger) *Store {
	s := Store{db: db}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}

func (s *Store) Get(ownerID int64) ([]Recipe, error) {
	rows, err := s.db.Query(`
		WITH owners_recipes AS (
		    SELECT id, name, owner_id 
			FROM recipes
			WHERE owner_id = $1
        )
		SELECT r.id, r.name, r.owner_id, re.id, re.food_item_id, re.amount
		FROM owners_recipes r
			LEFT JOIN recipe_entries re ON r.id = re.recipe_id`,
		ownerID,
	)
	if err != nil {
		return nil, err
	}
	recipes := make([]Recipe, 0)
	entries := make(map[int64][]Entry)
	lastRecipeID := int64(0)
	for rows.Next() {
		recipe := Recipe{}
		entry := Entry{}
		rows.Scan(&recipe.ID, &recipe.Name, &recipe.OwnerID, &entry.ID, &entry.FoodItemID, &entry.Amount)
		if lastRecipeID != recipe.ID {
			recipes = append(recipes, recipe)
			entries[recipe.ID] = make([]Entry, 0)
		}
		if entry.IsValid() {
			entries[recipe.ID] = append(entries[recipe.ID], entry)
		}
		lastRecipeID = recipe.ID
	}

	for i := 0; i < len(recipes); i++ {
		recipes[i].Entries = entries[recipes[i].ID]
	}

	return recipes, nil
}

func (s *Store) Add(recipe Recipe) (Recipe, error) {
	tx, err := s.db.Begin()

	err = tx.QueryRow(`
		INSERT INTO recipes AS r (name, owner_id)
		VALUES ($1, $2)
		RETURNING r.id`,
		recipe.Name,
		recipe.OwnerID,
	).Scan(&recipe.ID)
	if err != nil {
		return Recipe{}, err
	}

	for _, entry := range recipe.Entries {
		err = tx.QueryRow(`
			INSERT INTO recipe_entries AS re (amount, food_item_id, recipe_id)
			VALUES ($1, $2, $3)
			RETURNING re.id`,
			entry.Amount,
			entry.FoodItemIDOrNil(),
			recipe.ID,
		).Scan(&entry.ID)
	}
	err = tx.Commit()
	if err != nil {
		return Recipe{}, err
	}
	s.logger.Info("added new recipe", slog.Any("recipe", recipe))

	return recipe, nil
}

func (s *Store) Delete(id int64, ownerID int64) error {
	_, err := s.db.Query(`
		DELETE FROM recipes
		WHERE id = $1
		  AND owner_id = $2
	`, id, ownerID,
	)
	if err != nil {
		s.logger.Error("failed to delete recipe", slog.Int64("recipeID", id))
		return err
	}
	return nil
}
