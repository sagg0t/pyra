package nutrition

import (
	"context"

	"pyra/pkg/db"
)

type DishRef struct {
	UID DishUID
	Version DishVersion
}

type DishRepository interface {
	db.Repository[DishRepository]

	Index(context.Context) ([]Dish, error)
	FindByID(context.Context, DishID) (Dish, error)
	Versions(context.Context, DishUID) ([]Dish, error)
	FindAllByProductID(context.Context, ProductID) ([]Dish, error)
	FindAllByRefs(context.Context, []DishRef) ([]Dish, error)

	IsNameTaken(context.Context, DishName, DishUID) (bool, error)

	Create(context.Context, *Dish) error
}
