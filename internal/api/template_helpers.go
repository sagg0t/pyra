package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"strconv"

	"pyra/internal/api/routes"
)

var TemplateHelpers = template.FuncMap{
	"compactFloat":   compactFloat[float32],
	"compactFloat64": compactFloat[float64],
	"toJSON":         toJSON,
	"inputData":      inputData,

	// Food products URI helpers
	"editFoodProductURI": routes.EditFoodProduct,
	"foodProductURI":     routes.FoodProduct,

	// Dishes URI helpers
	"dishURI": routes.DishURI,
}

func compactFloat[T float32 | float64](f T) string {
	f64 := float64(f)
	if f64-float64(int64(f64)) == 0 {
		return strconv.FormatInt(int64(f64), 10)
	}

	return strconv.FormatFloat(f64, 'f', 2, 32)
}

func toJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func inputData(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("input data must have an even number of arguments")
	}

	m := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key := fmt.Sprint(values[i])
		m[key] = values[i+1]
	}

	return m, nil
}
