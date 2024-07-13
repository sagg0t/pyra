package foodproducts

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodProductsRepository interface {
	FindById(ctx context.Context, id uint64) (FoodProduct, error)
	Index(ctx context.Context) ([]FoodProduct, error)
	Create(ctx context.Context, params CreateRequest) (uint64, error)
	Delete(ctx context.Context, id uint64) error
}

type foodProductsRepo struct {
	db *pgxpool.Pool
}

func NewDB(db *pgxpool.Pool) FoodProductsRepository {
	return &foodProductsRepo{
		db: db,
	}
}

func (s *foodProductsRepo) FindById(ctx context.Context, id uint64) (FoodProduct, error) {
	row := s.db.QueryRow(ctx, "SELECT * FROM food_products WHERE id = $1 LIMIT 1", id)

	resultProduct := FoodProduct{}

	err := row.Scan(
		&resultProduct.ID,
		&resultProduct.Name,
		&resultProduct.Calories,
		&resultProduct.Proteins,
		&resultProduct.Fats,
		&resultProduct.Carbs,
		&resultProduct.CreatedAt,
		&resultProduct.UpdatedAt,
	)
	if err != nil {
		return FoodProduct{}, err
	}

	return resultProduct, nil
}

func (s *foodProductsRepo) Index(ctx context.Context) ([]FoodProduct, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM food_products;")
	if err != nil {
		return nil, err
	}

	foodProducts, err := pgx.CollectRows(rows, pgx.RowToStructByName[FoodProduct])
	if err != nil {
		return nil, err
	}

	return foodProducts, nil
}

func (s *foodProductsRepo) Create(ctx context.Context, params CreateRequest) (uint64, error) {
	params.Normalize()

	row := s.db.QueryRow(
		ctx,
		"INSERT INTO food_products (name, calories, proteins, fats, carbs) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		params.Name, params.Calories, params.Proteins, params.Fats, params.Carbs,
	)

	var newId uint64
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}

func (s *foodProductsRepo) Delete(ctx context.Context, id uint64) error {
	_, err := s.db.Exec(ctx, "DELETE FROM food_products WHERE id = $1", id)

	return err
}
