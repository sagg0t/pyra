// Package products
package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"pyra/pkg/db"
	"pyra/pkg/log"
	"pyra/pkg/nutrition"
)

type Repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) nutrition.ProductRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) BeginTx(ctx context.Context) (db.DBTX, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *Repository) WithTx(tx db.DBTX) nutrition.ProductRepository {
	return NewRepository(tx)
}

func (r *Repository) FindByID(
	ctx context.Context,
	id nutrition.ProductID,
) (nutrition.Product, error) {
	row := r.db.QueryRowContext(ctx, findByIDQuery, id)
	
	return r.scanProductRow(row)
}

func (r *Repository) FindByRef(
	ctx context.Context,
	uid nutrition.ProductUID,
	version nutrition.ProductVersion,
) (nutrition.Product, error) {
	row := r.db.QueryRowContext(ctx, findByRefQuery, uid, version)

	return r.scanProductRow(row)
}

func (r *Repository) Versions(
	ctx context.Context,
	uid nutrition.ProductUID,
) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, productVersionsQuery, uid)
	if err != nil {
		return nil, err
	}
	defer closeRows(ctx, rows)

	return r.scanProducts(rows)
}

func (r *Repository) FindAllByIDs(
	ctx context.Context,
	ids []nutrition.ProductID,
) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, findAllByIDsQuery, ids)
	if err != nil {
		return nil, err
	}
	defer closeRows(ctx, rows)

	return r.scanProducts(rows)
}

func (r *Repository) IsNameTaken(
	ctx context.Context,
	name nutrition.ProductName,
) (bool, error) {
	row := r.db.QueryRowContext(ctx, nameTakenQuery, name)
	if err := row.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	var one int
	if err := row.Scan(&one); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *Repository) ForDish(
	ctx context.Context,
	id nutrition.DishID,
) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, productsForDishQuery, id)
	if err != nil {
		return nil, err
	}
	defer closeRows(ctx, rows)

	return r.scanProducts(rows)
}

func (r *Repository) Index(ctx context.Context) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, indexProductsQuery)
	if err != nil {
		return nil, err
	}
	defer closeRows(ctx, rows)

	return r.scanProducts(rows)
}

func (r *Repository) Create(ctx context.Context, p *nutrition.Product) error {
	row := r.db.QueryRowContext(ctx, createProductQuery,
		p.UID, p.Name, p.Calories, p.Proteins, p.Fats, p.Carbs)

	return row.Scan(&p.ID, &p.Version)
}

func (r *Repository) CreateVersion(ctx context.Context, p *nutrition.Product) error {
	row := r.db.QueryRowContext(ctx, createProductVersionQuery,
		p.UID, p.Name, p.Calories, p.Proteins, p.Fats, p.Carbs)

	return row.Scan(&p.ID, &p.Version, &p.CreatedAt)
}

func (r *Repository) Delete(ctx context.Context, id nutrition.ProductID) error {
	_, err := r.db.ExecContext(ctx, deleteByIDQuery, id)

	return err
}

func (r *Repository) Update(ctx context.Context, product *nutrition.Product) error {
	result, err := r.db.ExecContext(ctx, updateProductQuery,
		product.UID, product.Version, product.Name,
		product.Calories, product.Proteins, product.Fats, product.Carbs)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d", rowsAffected)
	}

	return nil
}

func (r *Repository) Archive(ctx context.Context, id nutrition.ProductID, ts time.Time) error {
	result, err := r.db.ExecContext(ctx, "UPDATE products SET archived_at = $2 WHERE id = $1", id, ts)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d", rowsAffected)
	}

	return nil
}

func (r *Repository) Search(ctx context.Context, searchStr string) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, searchProductsQuery, searchStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	defer closeRows(ctx, rows)

	return r.scanProducts(rows)
}

func (r *Repository) MaxVersion(
	ctx context.Context,
	uid nutrition.ProductUID,
) (nutrition.ProductVersion, error) {
	row := r.db.QueryRowContext(ctx, maxProductVersionQuery, uid)

	var version nutrition.ProductVersion
	if err := row.Scan(&version); err != nil {
		return nutrition.ProductVersion(-1), err
	}

	return version, nil
}

func (r *Repository) UsedInDishes(ctx context.Context, id nutrition.ProductID) (bool, error) {
	row := r.db.QueryRowContext(ctx, usedInDishesQuery, id, nutrition.IngredientProduct)

	var one int
	err := row.Scan(&one)	

	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return one == 1, err
}

func (r *Repository) CountAll(ctx context.Context) (n int, err error) {
	row := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products;")
	err = row.Scan(&n)

	return
}

func (r *Repository) scanProducts(rows *sql.Rows) ([]nutrition.Product, error) {
	products := make([]nutrition.Product, 0)

	for rows.Next() {
		product, err := r.scanProductRows(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *Repository) scanProductRows(rows *sql.Rows) (nutrition.Product, error) {
	product := nutrition.Product{}

	err := rows.Scan(&product.ID, &product.UID, &product.Version, &product.Name,
		&product.Calories, &product.Proteins, &product.Fats, &product.Carbs,
		&product.CreatedAt, &product.UpdatedAt)

	return product, err
}

func (r *Repository) scanProductRow(row *sql.Row) (nutrition.Product, error) {
	product := nutrition.Product{}

	err := row.Scan(&product.ID, &product.UID, &product.Version, &product.Name,
		&product.Calories, &product.Proteins, &product.Fats, &product.Carbs,
		&product.CreatedAt, &product.UpdatedAt)

	return product, err
}

func closeRows(ctx context.Context, rows *sql.Rows) {
	if closeErr := rows.Close(); closeErr != nil {
		log.FromContext(ctx).WarnContext(ctx, "failed to close TX", "error", closeErr)
	}
}
