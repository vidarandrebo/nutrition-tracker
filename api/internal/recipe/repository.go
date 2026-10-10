package recipe

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"
)

type IRepository interface {
	Add(item TableRecipe) (TableRecipe, error)
	AddEntry(item TableRecipeFoodItemEntry) (TableRecipeFoodItemEntry, error)
	Delete(id int64) error
	Get(ownerID int64) ([]TableRecipe, error)
	GetByID(id int64) (TableRecipe, error)
	GetEntries(id int64) ([]TableRecipeFoodItemEntry, error)
}

type Repository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	r := Repository{db: db}
	r.logger = logger.With(slog.Any("module", reflect.TypeOf(r)))
	return &r
}

// Add returns the database or scan error if inserting the recipe fails,
// including sql.ErrNoRows when the insert returns no row.
func (r *Repository) Add(item TableRecipe) (TableRecipe, error) {
	scanErr := r.db.QueryRow(`
		INSERT INTO recipes AS r (name, owner_id)
		VALUES ($1, $2)
		RETURNING r.id`,
		item.Name,
		item.OwnerID,
	).Scan(&item.ID)

	if scanErr != nil {
		return TableRecipe{}, scanErr
	}

	r.logger.Info("added new recipe", slog.Any("recipe", item))

	return item, nil
}

// AddEntry returns the database or scan error if inserting the entry fails,
// including sql.ErrNoRows when the insert returns no row.
func (r *Repository) AddEntry(item TableRecipeFoodItemEntry) (TableRecipeFoodItemEntry, error) {
	scanErr := r.db.QueryRow(`
		INSERT INTO recipe_food_item_entries (amount, food_item_id, recipe_id)
		VALUES ($1, $2, $3)
		RETURNING id`,
		item.Amount,
		item.FoodItemID,
		item.RecipeID,
	).Scan(&item.ID)

	if scanErr != nil {
		return TableRecipeFoodItemEntry{}, scanErr
	}

	r.logger.Info("added new recipe entry", slog.Any("recipeEntry", item))

	return item, nil
}

// Delete returns the database error if deleting the recipe fails.
// It returns no error if the recipe does not exist.
func (r *Repository) Delete(id int64) error {
	r.logger.Info("deleting recipe", slog.Int64("id", id))
	_, err := r.db.Exec(`
		DELETE FROM recipes
		WHERE id = $1
	`, id,
	)
	if err != nil {
		return err
	}
	return nil
}

// Get returns database, scan, iteration, or close errors without partial
// results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no recipes match.
func (r *Repository) Get(ownerID int64) (recipes []TableRecipe, err error) {
	rows, queryErr := r.db.Query(`
		WITH owners_recipes AS (
		    SELECT id, name, owner_id 
			FROM recipes
			WHERE owner_id = $1
        )
		SELECT r.id, r.name, r.owner_id
		FROM owners_recipes r`,
		ownerID,
	)
	if queryErr != nil {
		return nil, queryErr
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			recipes, err = nil, errors.Join(err, closeErr)
		}
	}()
	recipes = make([]TableRecipe, 0)
	for rows.Next() {
		recipe := TableRecipe{}
		scanErr := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.OwnerID,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		recipes = append(recipes, recipe)
	}
	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return recipes, nil
}

// GetByID returns sql.ErrNoRows if the recipe does not exist, or the database
// or scan error if querying or scanning it fails.
func (r *Repository) GetByID(id int64) (TableRecipe, error) {
	recipe := TableRecipe{}
	scanErr := r.db.QueryRow(`
		SELECT id, name, owner_id
		FROM recipes 
		WHERE id = $1
	`, id).Scan(
		&recipe.ID,
		&recipe.Name,
		&recipe.OwnerID,
	)
	if scanErr != nil {
		return TableRecipe{}, scanErr
	}

	return recipe, nil
}

// GetEntries returns database, scan, iteration, or close errors without partial
// results. Concurrent errors during iteration and closing are joined.
// It returns an empty slice and no error when no entries match.
func (r *Repository) GetEntries(id int64) (entries []TableRecipeFoodItemEntry, err error) {
	rows, queryErr := r.db.Query(`
		SELECT id, amount, food_item_id, recipe_id
		FROM recipe_food_item_entries r
		WHERE recipe_id = $1`,
		id,
	)
	if queryErr != nil {
		return nil, queryErr
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			entries, err = nil, errors.Join(err, closeErr)
		}
	}()
	entries = make([]TableRecipeFoodItemEntry, 0)
	for rows.Next() {
		entry := TableRecipeFoodItemEntry{}
		scanErr := rows.Scan(
			&entry.ID,
			&entry.Amount,
			&entry.FoodItemID,
			&entry.RecipeID,
		)
		if scanErr != nil {
			return nil, scanErr
		}

		entries = append(entries, entry)
	}
	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	return entries, nil
}
