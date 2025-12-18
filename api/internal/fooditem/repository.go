package fooditem

import (
	"database/sql"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IRepository interface {
	Add(item TableFoodItem) (TableFoodItem, error)
	AddMicronutrient(item TableMicronutrient) (TableMicronutrient, error)
	AddPortionSize(item TablePortionSize) (TablePortionSize, error)
	Get(ownerID int64) ([]TableFoodItem, error)
	GetByID(id int64) (TableFoodItem, error)
	GetPortionSizes(foodItemID int64) ([]TablePortionSize, error)
	GetMicronutrients(foodItemID int64) ([]TableMicronutrient, error)
	Delete(id int64) error
	CheckOwnership(id int64, ownerID int64) error
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

func (r *Repository) AddMicronutrient(item TableMicronutrient) (TableMicronutrient, error) {
	err := r.db.QueryRow(`
			INSERT INTO food_item_micronutrients (name, amount, food_item_id) 
			VALUES ($1, $2, $3)`,
		item.Name,
		item.Amount,
		item.FoodItemID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add micronutrient to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TableMicronutrient{}, err
	}

	return item, nil
}

func (r *Repository) AddPortionSize(item TablePortionSize) (TablePortionSize, error) {
	err := r.db.QueryRow(`
			INSERT INTO food_item_portion_sizes (name, amount, food_item_id) 
			VALUES ($1, $2,$3)
    	`, item.Name, item.Amount, item.ID,
	).Scan(&item.ID)
	if err != nil {
		r.logger.Error("failed to add portion size to food item", slog.Int64("foodItemID", item.FoodItemID), slog.Any("err", err))
		return TablePortionSize{}, err
	}
	return item, nil
}

func (r *Repository) Get(ownerID int64) ([]TableFoodItem, error) {
	items := make([]TableFoodItem, 0)
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
		r.logger.Error("failed to query rows of food items", slog.Any("err", err))
		return nil, err
	}
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
			r.logger.Error("failed to scan rows of food items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

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
		r.logger.Error("failed to query rows of food items", slog.Any("err", err))
		return TableFoodItem{}, utils.ErrEntityNotFound
	}
	return item, nil
}

func (r *Repository) GetPortionSizes(foodItemID int64) ([]TablePortionSize, error) {
	items := make([]TablePortionSize, 0)
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
		r.logger.Error("failed to query rows of portion sizes", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TablePortionSize{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of portion sizes items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) GetMicronutrients(foodItemID int64) ([]TableMicronutrient, error) {
	items := make([]TableMicronutrient, 0)
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
		r.logger.Error("failed to query rows of micronutrients", slog.Any("err", err))
		return nil, err
	}
	for rows.Next() {
		item := TableMicronutrient{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Amount,
			&item.FoodItemID,
		)
		if err != nil {
			r.logger.Error("failed to scan rows of micronutrient items", slog.Any("err", err))
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *Repository) Delete(id int64) error {
	_, err := r.db.Query(`
		DELETE FROM food_items
		WHERE id = $1
	`, id,
	)
	if err != nil {
		r.logger.Error("failed to delete foodItem", slog.Int64("foodItemId", id))
		return err
	}
	return nil
}

func (r *Repository) CheckOwnership(id int64, ownerID int64) error {
	foodItem := TableFoodItem{}
	err := r.db.QueryRow(`
		SELECT id, owner_id 
		FROM food_items 
		WHERE id = $1
	`, id).Scan(
		&foodItem.ID,
		&foodItem.OwnerID,
	)
	if err != nil {
		return utils.ErrEntityNotFound
	}

	if foodItem.OwnerID == ownerID {
		return nil
	}

	return utils.ErrEntityNotOwned
}
