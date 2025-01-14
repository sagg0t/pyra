package dishes

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB: db,
	}
}

type Repository struct {
	DB *pgxpool.Pool
}

func (repo *Repository) Index(ctx context.Context) ([]Dish, error) {
	rows, err := repo.DB.Query(ctx, `SELECT
	dishes.*
	FROM dishes
	INNER JOIN ( SELECT DISTINCT
			uid,
			max(version) AS version
		FROM
			dishes
		GROUP BY
			uid) latest_dishes ON dishes.uid = latest_dishes.uid
	AND dishes.version = latest_dishes.version`)
	if err != nil {
		return nil, err
	}

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dish])
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (repo *Repository) FindByID(ctx context.Context, id uint64) (Dish, error) {
	row := repo.DB.QueryRow(ctx, "SELECT * FROM dishes WHERE id = $1 LIMIT 1", id)

	dish := Dish{}

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
		return Dish{}, err
	}

	return dish, nil
}

func (repo *Repository) Versions(ctx context.Context, uid string) ([]Dish, error) {
	rows, err := repo.DB.Query(ctx, "SELECT * FROM dishes WHERE uid = $1 ORDER BY version DESC LIMIT 20", uid)
	if err != nil {
		return nil, err
	}

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dish])
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (repo *Repository) FindAllByProductID(ctx context.Context, productID uint64) ([]Dish, error) {
	rows, err := repo.DB.Query(
		ctx,
		`SELECT dishes.*
		FROM dishes
		INNER JOIN dish_products ON dish_products.dish_id = dishes.id
		WHERE dish_products.food_product_id = $1;`,
		productID,
	)
	if err != nil {
		return nil, err
	}

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dish])
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
