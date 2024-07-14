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

func (fp *FoodProduct) Normalize(per float32) {
	// Normal values are per 100g
	ratio := 100.0 / per
	fp.Calories *= ratio
	fp.Proteins *= ratio
	fp.Fats *= ratio
	fp.Carbs *= ratio
}
