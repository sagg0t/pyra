package foodproducts

import "time"

type FoodProduct struct {
	ID uint64

	Name string

	Calories float32
	Proteins float32
	Fats     float32
	Carbs    float32

	CreatedAt time.Time
	UpdatedAt time.Time
}
