// Package products
package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"pyra/pkg/db"
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
	defer rows.Close()

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
	defer rows.Close()

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
	defer rows.Close()

	return r.scanProducts(rows)
}

func (r *Repository) Index(ctx context.Context) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, indexProductsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanProducts(rows)
}

func (r *Repository) Create(
	ctx context.Context,
	uid nutrition.ProductUID,
	name nutrition.ProductName,
	macro nutrition.Macro,
) (nutrition.Product, error) {
	row := r.db.QueryRowContext(ctx, createProductQuery,
		uid, name, macro.Calories, macro.Proteins, macro.Fats, macro.Carbs)

	product := nutrition.Product{
		UID:   uid,
		Name:  name,
		Macro: macro,
	}
	if err := row.Scan(&product.ID, &product.Version); err != nil {
		return nutrition.Product{}, err
	}

	return product, nil
}

func (r *Repository) CreateVersion(
	ctx context.Context,
	uid nutrition.ProductUID,
	name nutrition.ProductName,
	macro nutrition.Macro,
) (nutrition.Product, error) {
	row := r.db.QueryRowContext(ctx, createProductVersionQuery,
		uid, name, macro.Calories, macro.Proteins, macro.Fats, macro.Carbs)

	product := nutrition.Product{
		UID:   uid,
		Name:  name,
		Macro: macro,
	}

	if err := row.Scan(&product.ID, &product.Version, &product.CreatedAt); err != nil {
		return nutrition.Product{}, err
	}

	return product, nil
}

func (r *Repository) Delete(ctx context.Context, id nutrition.ProductID) error {
	_, err := r.db.ExecContext(ctx, deleteByIDQuery, id)

	return err
}

func (r *Repository) Update(
	ctx context.Context,
	id nutrition.ProductID,
	name nutrition.ProductName,
	macro nutrition.Macro,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, updateProductQuery,
		id, name,
		macro.Calories, macro.Proteins, macro.Fats, macro.Carbs)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			slog.ErrorContext(ctx, "error while rolling back a TX", "error", err)
		}

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			slog.ErrorContext(ctx, "error while rolling back a TX", "error", err)
		}
		return err
	}
	if rowsAffected != 1 {
		if err := tx.Rollback(); err != nil {
			slog.ErrorContext(ctx, "error while rolling back a TX", "error", err)
		}
		return fmt.Errorf("expected 1 row to be affected, got %d", rowsAffected)
	}

	return tx.Commit()
}

func (r *Repository) Search(
	ctx context.Context,
	searchStr string,
) ([]nutrition.Product, error) {
	rows, err := r.db.QueryContext(ctx, searchProductsQuery, searchStr)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

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

	err := rows.Scan(
		&product.ID,
		&product.UID,
		&product.Version,
		&product.Name,
		&product.Calories,
		&product.Proteins,
		&product.Fats,
		&product.Carbs,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nutrition.Product{}, err
	}

	return product, nil
}

func (r *Repository) scanProductRow(row *sql.Row) (nutrition.Product, error) {
	product := nutrition.Product{}

	err := row.Scan(
		&product.ID,
		&product.UID,
		&product.Version,
		&product.Name,
		&product.Calories,
		&product.Proteins,
		&product.Fats,
		&product.Carbs,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nutrition.Product{}, err
	}

	return product, nil
}
