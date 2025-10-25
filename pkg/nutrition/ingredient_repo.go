package nutrition

import (
	"context"
	"pyra/pkg/db"
)

type IngredientRepository interface {
	db.Repository[IngredientRepository]

	GetIngredientables(context.Context, []IngredientInfo) ([]Ingredientable, error)

	CreateIngredients(context.Context, []Ingredient) error
}
