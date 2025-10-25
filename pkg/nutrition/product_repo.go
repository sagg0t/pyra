package nutrition

import (
	"context"
	"time"

	"pyra/pkg/db"
)

type ProductRepository interface {
	db.Repository[ProductRepository]

	Index(context.Context) ([]Product, error)
	FindAllByIDs(context.Context, []ProductID) ([]Product, error)
	ForDish(context.Context, DishID) ([]Product, error)

	FindByID(context.Context, ProductID) (Product, error)
	FindByRef(context.Context, ProductUID, ProductVersion) (Product, error)

	Versions(context.Context, ProductUID) ([]Product, error)

	Create(context.Context, *Product) error
	CreateVersion(context.Context, *Product) error
	Delete(context.Context, ProductID) error
	Update(context.Context, *Product) error
	Archive(context.Context, ProductID, time.Time) error

	CountAll(context.Context) (int, error)

	IsNameTaken(context.Context, ProductName) (bool, error)
	UsedInDishes(context.Context, ProductID) (bool, error)
	MaxVersion(context.Context, ProductUID) (ProductVersion, error)

	Search(ctx context.Context, searchStr string) ([]Product, error)
}
