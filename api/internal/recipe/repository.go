package recipe

import (
	"database/sql"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IRepository interface {
	Add(item TableRecipe) (TableRecipe, error)
	AddEntry(item TableRecipeEntry) (TableRecipeEntry, error)
	Delete(id int64) error
	Get(ownerID int64) ([]TableRecipe, error)
	GetByID(id int64) (TableRecipe, error)
	GetEntries(id int64) ([]TableRecipeEntry, error)
	CheckOwnership(id int64, ownerID int64) error
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

func (r *Repository) Add(item TableRecipe) (TableRecipe, error) {
	scanErr := r.db.QueryRow(`
		INSERT INTO recipes AS r (name, owner_id)
		VALUES ($1, $2)
		RETURNING r.id`,
		item.Name,
		item.OwnerID,
	).Scan(&item.ID)

	if scanErr != nil {
		r.logger.Error("failed to add recipe", slog.Any("err", scanErr), slog.Int64("userID", item.OwnerID))
		return TableRecipe{}, utils.ErrUnknown
	}

	r.logger.Info("added new recipe", slog.Any("recipe", item))

	return item, nil
}

func (r *Repository) AddEntry(item TableRecipeEntry) (TableRecipeEntry, error) {
	scanErr := r.db.QueryRow(`
		INSERT INTO recipe_entries (amount, food_item_id, recipe_id)
		VALUES ($1, $2, $3)
		RETURNING id`,
		item.Amount,
		item.FoodItemID,
		item.RecipeID,
	).Scan(&item.ID)

	if scanErr != nil {
		r.logger.Error("failed to add recipe entry", slog.Any("err", scanErr), slog.Int64("recipeID", item.RecipeID))
		return TableRecipeEntry{}, utils.ErrUnknown
	}

	r.logger.Info("added new recipe entry", slog.Any("recipeEntry", item))

	return item, nil
}

func (r *Repository) Delete(id int64) error {
	r.logger.Info("deleting recipe", slog.Int64("id", id))
	_, err := r.db.Query(`
		DELETE FROM recipes
		WHERE id = $1
	`, id,
	)
	if err != nil {
		r.logger.Error("failed to delete recipe", slog.Int64("recipeID", id))
		return err
	}
	return nil
}

func (r *Repository) Get(ownerID int64) ([]TableRecipe, error) {
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
		r.logger.Error("failed to query rows of recipes", slog.Any("err", queryErr))
		return nil, utils.ErrUnknown
	}
	recipes := make([]TableRecipe, 0)
	for rows.Next() {
		recipe := TableRecipe{}
		scanErr := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.OwnerID,
		)
		if scanErr != nil {
			r.logger.Error("failed to scan rows of recipes", slog.Any("err", scanErr))
			return nil, utils.ErrUnknown
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

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
		return TableRecipe{}, utils.ErrEntityNotFound
	}

	return recipe, nil
}

func (r *Repository) GetEntries(id int64) ([]TableRecipeEntry, error) {
	rows, queryErr := r.db.Query(`
		SELECT id, amount, food_item_id, recipe_id
		FROM recipe_entries r
		WHERE recipe_id = $1`,
		id,
	)
	if queryErr != nil {
		r.logger.Error("failed to query rows of recipe entries", slog.Any("err", queryErr))
		return nil, utils.ErrUnknown
	}
	entries := make([]TableRecipeEntry, 0)
	for rows.Next() {
		entry := TableRecipeEntry{}
		scanErr := rows.Scan(
			&entry.ID,
			&entry.Amount,
			&entry.FoodItemID,
			&entry.RecipeID,
		)
		if scanErr != nil {
			r.logger.Error("failed to scan rows of recipe entries", slog.Any("err", scanErr))
			return nil, utils.ErrUnknown
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func (r *Repository) CheckOwnership(id int64, ownerID int64) error {
	recipe := TableRecipe{}
	err := r.db.QueryRow(`
		SELECT id, owner_id 
		FROM recipes 
		WHERE id = $1
	`, id).Scan(
		&recipe.ID,
		&recipe.OwnerID,
	)
	if err != nil {
		return utils.ErrEntityNotFound
	}

	if recipe.OwnerID == ownerID {
		return nil
	}

	return utils.ErrEntityNotOwned
}
