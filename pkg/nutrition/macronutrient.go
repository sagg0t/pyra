package nutrition

import (
	"errors"
	"fmt"
)

var ErrMacroNegative = errors.New("value of macronutrient can't be negative")

type Macro struct {
	Calories float32
	Proteins float32
	Fats     float32
	Carbs    float32
}

type MacroErrors struct {
	Calories error
	Proteins error
	Fats     error
	Carbs    error
}

func (e *MacroErrors) HasErrors() bool {
	caloriesErr := e.Calories != nil
	proteinsErr := e.Proteins != nil
	fatsErr := e.Fats != nil
	carbsErr := e.Carbs != nil

	return caloriesErr || proteinsErr || fatsErr || carbsErr
}

const macroErrFmt = `calories: %w
proteins: %w
fats: %w
carbs: %w`

func (e *MacroErrors) Error() string {
	return fmt.Errorf(macroErrFmt,
		e.Calories, e.Proteins, e.Fats, e.Carbs,
	).Error()
}

func NewMacro(calories, proteins, fats, carbs float32) (Macro, MacroErrors) {
	mErrors := MacroErrors{}

	if calories < 0 {
		mErrors.Calories = ErrMacroNegative
	}

	if proteins < 0 {
		mErrors.Proteins = ErrMacroNegative
	}

	if fats < 0 {
		mErrors.Fats = ErrMacroNegative
	}

	if carbs < 0 {
		mErrors.Carbs = ErrMacroNegative
	}

	return Macro{
		Calories: calories,
		Proteins: proteins,
		Fats:     fats,
		Carbs:    carbs,
	}, mErrors
}

func (m *Macro) Normalize(per float32) {
	ratio := per / 100.0

	m.Calories *= ratio
	m.Proteins *= ratio
	m.Fats *= ratio
	m.Carbs *= ratio
}
