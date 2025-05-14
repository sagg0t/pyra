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
	ProductID

	Amount float32
	Unit   MeasurementUnit
}

type MeasurementUnit int32

const (
	InvalidUnit MeasurementUnit = iota
	Gramms
	Milliliter
	Unit
)

func NewMeasurementUnit(iunit int32) MeasurementUnit {
	switch iunit {
	case 1:
		return Gramms
	case 2:
		return Milliliter
	case 3:
		return Unit
	default:
		return InvalidUnit
	}
}

func NewIngredient(dishID uint64, productID uint64, amt float32, unit int32) (Ingredient, error) {
	mUnit := NewMeasurementUnit(unit)
	if mUnit == InvalidUnit {
		return Ingredient{}, fmt.Errorf("%w: %v", ErrInvalidUnit, unit)
	}

	if amt < 0 {
		return Ingredient{}, ErrNegativeAmount
	}

	return Ingredient{
		DishID:    DishID(dishID),
		ProductID: ProductID(productID),
		Amount:    amt,
		Unit:      mUnit,
	}, nil
}
