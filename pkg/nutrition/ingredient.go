package nutrition

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidUnit    = errors.New("invalid value for a measurement unit")
	ErrNegativeAmount = errors.New("ingredient amount can't be less than 0")
)

type Ingredient struct {
	DishID

	IngredientableID   uint64
	IngredientableType IngredientableType

	Amount float32
	Unit   MeasurementUnit

	Idx uint16
}

type IngredientErrors struct {
	DishID error

	IngredientableID   error
	IngredientableType error

	Amount error
	Unit   error

	Idx error
}

func NewIngredient(
	dishID DishID,
	ingredientableID uint64,
	ingredientableType IngredientableType,
	amt float32,
	unit MeasurementUnit,
) (Ingredient, error) {
	if unit == InvalidUnit {
		return Ingredient{}, fmt.Errorf("%w: %v", ErrInvalidUnit, unit)
	}

	if amt < 0 {
		return Ingredient{}, ErrNegativeAmount
	}

	return Ingredient{
		DishID:           dishID,
		IngredientableID: ingredientableID,
		Amount:           amt,
		Unit:             unit,
	}, nil
}

type IngredientableType uint16

const (
	IngredientNone IngredientableType = iota
	IngredientProduct
	IngredientDish
)

func (it IngredientableType) String() string {
	switch it {
	case IngredientNone:
		return "IngredientNone"
	case IngredientProduct:
		return "IngredientProduct"
	case IngredientDish:
		return "IngredientDish"
	default:
		panic("WTF?")
	}
}

type Ingredientable struct {
	ID   uint64

	Macro

	Info IngredientInfo
}

const ingredientableDebugFormat = `
	ID: %d
	Calories: %.2f
	Proteins: %.2f
	Fats: %.2f
	Carbs: %.2f
	Info:
		Idx: %d
		UID: %s
		Version: %d
		Type: %s
		Amount: %.2f
		Unit: %s

`
func (ing Ingredientable) String() string {
	return fmt.Sprintf(ingredientableDebugFormat, ing.ID,
		ing.Calories.Float(), ing.Proteins.Float(), ing.Fats.Float(), ing.Carbs.Float(),
		ing.Info.Idx, ing.Info.UID, ing.Info.Version, ing.Info.Type,
		ing.Info.Amount, ing.Info.Unit)
}
