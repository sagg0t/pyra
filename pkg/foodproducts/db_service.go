package foodproducts

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodProductsService interface {
	FindById(ctx context.Context, id uint64) (FoodProduct, error)
	Index(ctx context.Context) ([]FoodProduct, error)
}

type service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) FoodProductsService {
	return &service{
		db: db,
	}
}

func (s *service) FindById(ctx context.Context, id uint64) (FoodProduct, error) {
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

func (s *service) Index(ctx context.Context) ([]FoodProduct, error) {
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

func (s *service) Create(ctx context.Context, product FoodProduct) error {
	_, err := s.db.Exec(
		ctx,
		"INSERT INTO food_products (name, calories, proteins, fats, carbs) VALUES ($1, $2, $3, $4, $5)",
		product.Name, product.Calories, product.Proteins, product.Fats, product.Carbs,
	)

	return err
}
