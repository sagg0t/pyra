package dishes

import (
	"time"
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

	// Ingredients []foodproducts.FoodProduct

	CreatedAt time.Time
	UpdatedAt time.Time
}
