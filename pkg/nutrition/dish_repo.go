package nutrition

import "context"

type DishRepository interface {
	Index(ctx context.Context) ([]Dish, error)
	FindByID(ctx context.Context, id DishID) (Dish, error)
	Versions(ctx context.Context, uid DishUID) ([]Dish, error)
	FindAllByProductID(ctx context.Context, productID ProductID) ([]Dish, error)
}
