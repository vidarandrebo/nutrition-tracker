package fooditem

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type IRepository interface {
	Add(item TableFoodItem) (TableFoodItem, error)
	AddMicronutrient(item TableFoodItemMacronutrient) (TableFoodItemMacronutrient, error)
	AddPortionSize(item TableFoodItemPortionSize) (TableFoodItemPortionSize, error)
	Get(ownerID int64) ([]TableFoodItem, error)
	GetByID(id int64) (TableFoodItem, error)
	GetPortionSizes(foodItemID int64) ([]TableFoodItemPortionSize, error)
	GetMicronutrients(foodItemID int64) ([]TableFoodItemMacronutrient, error)
	Delete(id int64) error
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

// Add returns sql.ErrNoRows if the insert returns no row, or the database error
// if inserting or scanning the new food item fails.
func (r *Repository) Add(item TableFoodItem) (TableFoodItem, error) {
	err := r.db.QueryRow(`
		INSERT INTO food_items AS fi (manufacturer, product, protein, carbohydrate, fat, kcal, public, source, owner_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING fi.id`,
		item.Manufacturer,
		item.Product,
		item.Protein,
		item.Carbohydrate,
		item.Fat,
		item.KCal,
		item.Public,
		item.Source,
		item.OwnerID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add food item", slog.Int64("userID", item.OwnerID), slog.Any("err", err))
		return TableFoodItem{}, err
	}
	r.logger.Info("added new food item", slog.Any("food item", item))

	return item, nil
}

// AddMicronutrient returns sql.ErrNoRows if the insert returns no row, or the
// database error if inserting or scanning the new micronutrient fails.
func (r *Repository) AddMicronutrient(item TableFoodItemMacronutrient) (TableFoodItemMacronutrient, error) {
	err := r.db.QueryRow(`
			INSERT INTO food_item_micronutrients (name, amount, food_item_id) 
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		item.Name,
		item.Amount,
		item.FoodItemID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add micronutrient to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TableFoodItemMacronutrient{}, err
	}

	return item, nil
}

// AddPortionSize returns sql.ErrNoRows if the insert returns no row, or the
// database error if inserting or scanning the new portion size fails.
func (r *Repository) AddPortionSize(item TableFoodItemPortionSize) (TableFoodItemPortionSize, error) {
	err := r.db.QueryRow(`
			INSERT INTO food_item_portion_sizes (name, amount, food_item_id) 
			VALUES ($1, $2,$3)
			RETURNING id
    	`, item.Name, item.Amount, item.ID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add portion size to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TableFoodItemPortionSize{}, err
	}
	return item, nil
}

// Get returns the database error if querying, scanning, iterating or closing food items fails.
// It returns an empty slice and no error when no food items match.
func (r *Repository) Get(ownerID int64) (items []TableFoodItem, err error) {
	items = make([]TableFoodItem, 0)
	rows, err := r.db.Query(`
		WITH owned_fi AS (
		    SELECT *
			FROM food_items
			WHERE public = TRUE 
 			  OR owner_id = $1
        )		
		SELECT fi.id, fi.manufacturer, fi.product, fi.protein, fi.carbohydrate, fi.fat, fi.kcal, fi.public, fi.source, fi.owner_id
		FROM owned_fi fi
		`,
		ownerID,
	)
	if err != nil {
		r.logger.Error("failed to query rows of food items", slog.Int64("ownerID", ownerID), slog.Any("err", err))
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error("failed to close food item rows", slog.Int64("ownerID", ownerID), slog.Any("err", closeErr))
			items, err = nil, closeErr
		}
	}()
	for rows.Next() {
		item := TableFoodItem{}
		err = rows.Scan(
			&item.ID,
			&item.Manufacturer,
			&item.Product,
			&item.Protein,
			&item.Carbohydrate,
			&item.Fat,
			&item.KCal,
			&item.Public,
			&item.Source,
			&item.OwnerID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of food items", slog.Int64("ownerID", ownerID), slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("failed to iterate food items", slog.Int64("ownerID", ownerID), slog.Any("err", err))
		return nil, err
	}
	return items, nil
}

// GetByID returns sql.ErrNoRows if the food item does not exist, or the
// database error if querying or scanning it fails for another reason.
func (r *Repository) GetByID(id int64) (TableFoodItem, error) {
	item := TableFoodItem{}
	err := r.db.QueryRow(`
		SELECT fi.id, fi.manufacturer, fi.product, fi.protein, fi.carbohydrate, fi.fat, fi.kcal, fi.public, fi.source, fi.owner_id
		FROM food_items fi
		WHERE fi.id = $1
		`,
		id,
	).Scan(&item.ID, &item.Manufacturer, &item.Product, &item.Protein, &item.Carbohydrate, &item.Fat, &item.KCal, &item.Public, &item.Source, &item.OwnerID)
	if err != nil {
		r.logger.Error("failed to query food item", slog.Int64("foodItemID", id), slog.Any("err", err))
		return TableFoodItem{}, err
	}
	return item, nil
}

// GetPortionSizes returns the database error if querying, scanning, iterating or closing portion sizes fails.
// It returns an empty slice and no error when no portion sizes match.
func (r *Repository) GetPortionSizes(foodItemID int64) (items []TableFoodItemPortionSize, err error) {
	items = make([]TableFoodItemPortionSize, 0)
	rows, err := r.db.Query(`
		WITH fi_portions AS (
		    SELECT *
			FROM food_item_portion_sizes
			WHERE food_item_id = $1
        )		
		SELECT ps.id, ps.name, ps.amount, ps.food_item_id
		FROM fi_portions ps
		`,
		foodItemID,
	)
	if err != nil {
		r.logger.Error("failed to query rows of portion sizes", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error("failed to close portion size rows", slog.Int64("foodItemID", foodItemID), slog.Any("err", closeErr))
			items, err = nil, closeErr
		}
	}()
	for rows.Next() {
		item := TableFoodItemPortionSize{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of portion sizes items", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("failed to iterate portion sizes", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, err
	}
	return items, nil
}

// GetMicronutrients returns the database error if querying, scanning, iterating or closing micronutrients fails.
// It returns an empty slice and no error when no micronutrients match.
func (r *Repository) GetMicronutrients(foodItemID int64) (items []TableFoodItemMacronutrient, err error) {
	items = make([]TableFoodItemMacronutrient, 0)
	rows, err := r.db.Query(`
		WITH fi_micronutrients AS (
		    SELECT *
			FROM food_item_micronutrients
			WHERE food_item_id = $1
        )		
		SELECT mn.id, mn.name, mn.amount, mn.food_item_id
		FROM fi_micronutrients mn
		`,
		foodItemID,
	)
	if err != nil {
		r.logger.Error("failed to query rows of micronutrients", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			r.logger.Error("failed to close micronutrient rows", slog.Int64("foodItemID", foodItemID), slog.Any("err", closeErr))
			items, err = nil, closeErr
		}
	}()
	for rows.Next() {
		item := TableFoodItemMacronutrient{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of micronutrient items", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("failed to iterate micronutrients", slog.Int64("foodItemID", foodItemID), slog.Any("err", err))
		return nil, err
	}
	return items, nil
}

// Delete returns the database error if deleting the food item fails.
// It returns no error when the food item does not exist.
func (r *Repository) Delete(id int64) error {
	_, err := r.db.Exec(`
		DELETE FROM food_items
		WHERE id = $1
	`, id,
	)
	if err != nil {
		r.logger.Error("failed to delete food item", slog.Int64("foodItemID", id), slog.Any("err", err))
		return err
	}
	return nil
}
