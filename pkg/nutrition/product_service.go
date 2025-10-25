package nutrition

import (
	"context"
	"errors"
	"time"

	"pyra/pkg/db"

	"github.com/google/uuid"
)

var (
	ErrProductNameTaken = errors.New("product with such name already exists")
	ErrProductInvalid   = errors.New("product validation failed")
	ErrNotLatest        = errors.New("can only update the latest version")
	ErrProductUsed      = errors.New("cannot modify product that has been used")
	ErrArchived         = errors.New("cannot modify archived product")
)

func ListProducts(ctx context.Context, repo ProductRepository) ([]Product, error) {
	return repo.Index(ctx)
}

func FindProductByID(ctx context.Context, repo ProductRepository, id ProductID) (Product, error) {
	return repo.FindByID(ctx, id)
}

func CreateProduct(ctx context.Context, repo ProductRepository, product *Product) (err error) {
	tx, err := repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer db.RollbackGuard(ctx, tx, &err)

	repo = repo.WithTx(tx)

	product.UID = ProductUID(uuid.New().String())

	isTaken, err := repo.IsNameTaken(ctx, product.Name)
	if err != nil {
		return err
	} else if isTaken {
		product.Errors.Name = ErrProductNameTaken
		return ErrProductInvalid
	}

	err = repo.Create(ctx, product)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateProduct(ctx context.Context, repo ProductRepository, product *Product) (err error) {
	tx, err := repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer db.RollbackGuard(ctx, tx, &err)

	repo = repo.WithTx(tx)

	currentState, err := repo.FindByRef(ctx, product.UID, product.Version)
	if err != nil {
		return err
	}

	if currentState.IsArchived() {
		return ErrArchived
	}

	// FIX: actually being used in dishes is not a problem. Only if that dish has been
	// used in the menu - then it becomes untouchable.
	usedInDishes, err := repo.UsedInDishes(ctx, currentState.ID)
	if err != nil {
		return err
	}

	if usedInDishes {
		archivedAt := time.Now().UTC()
		err := repo.Archive(ctx, product.ID, archivedAt)
		if err != nil {
			return err
		}

		err = repo.CreateVersion(ctx, product)
		if err != nil {
			return err
		}
	} else {
		err := repo.Update(ctx, product)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func DeleteProduct(
	ctx context.Context,
	repo ProductRepository,
	uid ProductUID,
	version ProductVersion,
) error {
	tx, err := repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer db.RollbackGuard(ctx, tx, &err)

	repo = repo.WithTx(tx)

	product, err := repo.FindByRef(ctx, uid, version)
	if err != nil {
		return err
	}

	if product.IsArchived() {
		return ErrArchived
	}

	usedInDishes, err := repo.UsedInDishes(ctx, product.ID)
	if err != nil {
		return err
	}

	if usedInDishes {
		return ErrProductUsed
	}

	return tx.Commit()
}
