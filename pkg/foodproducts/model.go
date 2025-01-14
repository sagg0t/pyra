package foodproducts

import (
	"time"
)

type FoodProduct struct {
	ID      uint64 `json:"id"`
	UID     string `json:"uid"`
	Version int32  `json:"version"`

	Name string `json:"name"`

	Calories float32 `json:"calories"`
	Proteins float32 `json:"proteins"`
	Fats     float32 `json:"fats"`
	Carbs    float32 `json:"carbs"`

	CreatedAt time.Time `json:""`
	UpdatedAt time.Time `json:""`
}

func (fp *FoodProduct) Normalize(per float32) {
	// Normal values are per 100g
	ratio := 100.0 / per
	fp.Calories *= ratio
	fp.Proteins *= ratio
	fp.Fats *= ratio
	fp.Carbs *= ratio
}

type FoodProductAttributes struct {
	Name string `json:"name"`

	Calories float32 `json:"calories"`
	Proteins float32 `json:"proteins"`
	Fats     float32 `json:"fats"`
	Carbs    float32 `json:"carbs"`
}
