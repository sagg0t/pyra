package dishes

import (
	"time"

	"pyra/pkg/dishproducts"
)

type Dish struct {
	ID      uint64
	UID     string
	Version int32

	Name string

	Calories float32
	Proteins float32
	Fats     float32
	Carbs    float32

	CreatedAt time.Time
	UpdatedAt time.Time

	ingredients []dishproducts.DishProduct
}

func (dish *Dish) Ingredients() []dishproducts.DishProduct {
	return dish.ingredients
}

func (dish *Dish) SetIngredients(ingredients []dishproducts.DishProduct) {
	dish.ingredients = ingredients
}
