package pyra

import (
	"time"
)

type Base struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Ingredient struct {
	Base
	Macros
	Name string `json:"name"`
}

// A single dish, a drink or a snack
type MenuItem struct {
	Base
	Macros
	Name string `json:"name"`
}

type MealSetType int

const (
	BreakfastType MealSetType = iota
	DinnerType
)

// A set of dishes - either a breakfast or a dinner.
type MealSet struct {
	Base
	Type  MealSetType
	Items []MenuItem
}

type Macros struct {
	Proteins float64 `json:"proteins"`
	Fats     float64 `json:"fats"`
	Carbs    float64 `json:"carbohydrates"`
	Calories float64 `json:"calories"`
}
