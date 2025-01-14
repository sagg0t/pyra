package foodproducts

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FoodProductsRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *FoodProductsRepository {
	return &FoodProductsRepository{
		db: db,
	}
}

func (s *FoodProductsRepository) FindById(ctx context.Context, id uint64) (FoodProduct, error) {
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
		&resultProduct.UID,
		&resultProduct.Version,
	)
	if err != nil {
		return FoodProduct{}, err
	}

	return resultProduct, nil
}

func (s *FoodProductsRepository) Versions(ctx context.Context, uid string) ([]FoodProduct, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM food_products WHERE uid = $1", uid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[FoodProduct])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *FoodProductsRepository) FindAllByIds(ctx context.Context, ids []uint64) ([]FoodProduct, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM food_products WHERE id in $0", ids)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[FoodProduct])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *FoodProductsRepository) ForDish(ctx context.Context, id uint64) ([]FoodProduct, error) {
	rows, err := s.db.Query(
		ctx,
		`SELECT food_products.*
		FROM dish_products
		JOIN food_products ON dish_products.food_product_id = food_products.id
		WHERE dish_products.dish_id = $1;`,
		id,
	)
	if err != nil {
		return nil, err
	}

	foodProducts, err := pgx.CollectRows(rows, pgx.RowToStructByName[FoodProduct])
	if err != nil {
		return nil, err
	}

	return foodProducts, nil
}

func (s *FoodProductsRepository) Index(ctx context.Context) ([]FoodProduct, error) {
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

func (s *FoodProductsRepository) Create(ctx context.Context, params FoodProduct) (uint64, error) {
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

func (s *FoodProductsRepository) Delete(ctx context.Context, id uint64) error {
	_, err := s.db.Exec(ctx, "DELETE FROM food_products WHERE id = $1", id)

	return err
}

func (s *FoodProductsRepository) Update(ctx context.Context, product FoodProduct) error {
	_, err := s.db.Exec(
		ctx,
		"UPDATE food_products SET name = $2, calories = $3, proteins = $4, fats = $5, carbs = $6 WHERE id = $1",
		product.ID, product.Name, product.Calories, product.Proteins, product.Fats, product.Carbs,
	)

	return err
}

func (s *FoodProductsRepository) Search(ctx context.Context, searchStr string) ([]FoodProduct, error) {
	rows, err := s.db.Query(
		ctx,
		"SELECT id, name, calories, proteins, fats, carbs FROM food_products WHERE name ILIKE '%' || $1 || '%'",
		searchStr,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	products := make([]FoodProduct, 0)
	for rows.Next() {
		var product FoodProduct

		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Calories,
			&product.Proteins,
			&product.Fats,
			&product.Carbs,
		)

		products = append(products, product)
	}

	return products, nil
}
