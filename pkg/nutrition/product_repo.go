package nutrition

import "context"

type ProductRepository interface {
	Index(context.Context) ([]Product, error)
	FindAllByIDs(context.Context, []ProductID) ([]Product, error)
	ForDish(context.Context, DishID) ([]Product, error)

	FindByID(context.Context, ProductID) (Product, error)
	Create(context.Context, ProductUID, ProductName, Macro) (Product, error)
	CreateVersion(context.Context, ProductUID, ProductName, Macro) (Product, error)
	Delete(context.Context, ProductID) error
	Update(context.Context, ProductID, ProductName, Macro) error

	IsNameTaken(context.Context, ProductName) (bool, error)

	Versions(context.Context, ProductUID) ([]Product, error)

	Search(ctx context.Context, searchStr string) ([]Product, error)
}
