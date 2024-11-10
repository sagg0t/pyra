package dishes

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Index(ctx context.Context) ([]Dish, error)
	FindByID(ctx context.Context, id uint64) (Dish, error)
	Versions(ctx context.Context, uid string) ([]Dish, error)
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &dishesRepo{
		db: db,
	}
}

type dishesRepo struct {
	db *pgxpool.Pool
}

func (repo *dishesRepo) Index(ctx context.Context) ([]Dish, error) {
	rows, err := repo.db.Query(ctx, `SELECT
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

func (repo *dishesRepo) FindByID(ctx context.Context, id uint64) (Dish, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM dishes WHERE id = $1 LIMIT 1", id)

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

func (repo *dishesRepo) Versions(ctx context.Context, uid string) ([]Dish, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM dishes WHERE uid = $1 ORDER BY version DESC LIMIT 20", uid)
	if err != nil {
		return nil, err
	}

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[Dish])
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
