package dishes

import (
	"context"
	"database/sql"

	"pyra/pkg/db"
	"pyra/pkg/nutrition"
)

func NewRepository(db db.DBTX) *Repository {
	return &Repository{
		db: db,
	}
}

type Repository struct {
	db db.DBTX
}

const listDishesQuery = `SELECT
	dishes.*
	FROM dishes
	INNER JOIN ( SELECT DISTINCT
			uid,
			max(version) AS version
		FROM
			dishes
		GROUP BY
			uid) latest_dishes ON dishes.uid = latest_dishes.uid
	AND dishes.version = latest_dishes.version`

func (repo *Repository) Index(ctx context.Context) ([]nutrition.Dish, error) {
	rows, err := repo.db.QueryContext(ctx, listDishesQuery)
	if err != nil {
		return nil, err
	}

	dishes, err := repo.scanAllRows(rows)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (repo *Repository) FindByID(ctx context.Context, id nutrition.DishID) (nutrition.Dish, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT * FROM dishes WHERE id = $1 LIMIT 1", id)

	dish, err := repo.scanRow(row)
	if err != nil {
		return nutrition.Dish{}, err
	}

	return dish, nil
}

func (repo *Repository) Versions(ctx context.Context, uid nutrition.DishUID) ([]nutrition.Dish, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM dishes WHERE uid = $1 ORDER BY version DESC LIMIT 20", uid)
	if err != nil {
		return nil, err
	}

	dishes, err := repo.scanAllRows(rows)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

const dishesByProductQuery = `SELECT dishes.*
		FROM dishes
		INNER JOIN dish_products ON dish_products.dish_id = dishes.id
		WHERE dish_products.food_product_id = $1;`

func (repo *Repository) FindAllByProductID(ctx context.Context, productID nutrition.ProductID) ([]nutrition.Dish, error) {
	rows, err := repo.db.QueryContext(ctx, dishesByProductQuery, productID)
	if err != nil {
		return nil, err
	}

	dishes, err := repo.scanAllRows(rows)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *Repository) scanRow(row *sql.Row) (nutrition.Dish, error) {
	dish := nutrition.Dish{}

	err := row.Scan(
		&dish.ID,
		&dish.UID,
		&dish.Version,
		&dish.Name,
		&dish.Calories,
		&dish.Proteins,
		&dish.Fats,
		&dish.Carbs,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)
	if err != nil {
		return nutrition.Dish{}, err
	}

	return dish, nil
}

func (r *Repository) scanRows(rows *sql.Rows) (nutrition.Dish, error) {
	dish := nutrition.Dish{}

	err := rows.Scan(
		&dish.ID,
		&dish.UID,
		&dish.Version,
		&dish.Name,
		&dish.Calories,
		&dish.Proteins,
		&dish.Fats,
		&dish.Carbs,
		&dish.CreatedAt,
		&dish.UpdatedAt,
	)
	if err != nil {
		return nutrition.Dish{}, err
	}

	return dish, nil
}

func (r *Repository) scanAllRows(rows *sql.Rows) ([]nutrition.Dish, error) {
	dishes := make([]nutrition.Dish, 0)

	for rows.Next() {
		dish := nutrition.Dish{}

		err := rows.Scan(
			&dish.ID,
			&dish.UID,
			&dish.Version,
			&dish.Name,
			&dish.Calories,
			&dish.Proteins,
			&dish.Fats,
			&dish.Carbs,
			&dish.CreatedAt,
			&dish.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}
