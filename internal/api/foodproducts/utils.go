package foodproducts

import (
	"net/http"
	"strconv"

	"pyra/pkg/foodproducts"
)

func productID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}

func paramsFromForm(fetch func(key string) string) (ProductForm, error) {
	form := ProductForm{
		FoodProduct: foodproducts.FoodProduct{
			Name: fetch("name"),
		},
		Errors: map[string]string{},
	}

	calories64, err := strconv.ParseFloat(fetch("calories"), 32)
	if err != nil {
		return form, err
	}
	form.Calories = float32(calories64)

	proteins64, err := strconv.ParseFloat(fetch("proteins"), 32)
	if err != nil {
		return form, err
	}
	form.Proteins = float32(proteins64)

	fats64, err := strconv.ParseFloat(fetch("fats"), 32)
	if err != nil {
		return form, err
	}
	form.Fats = float32(fats64)

	carbs64, err := strconv.ParseFloat(fetch("carbs"), 32)
	if err != nil {
		return form, err
	}
	form.Carbs = float32(carbs64)

	per64, err := strconv.ParseFloat(fetch("per"), 32)
	if err != nil {
		return form, err
	}
	form.Per = float32(per64)

	return form, nil
}

const (
	errNoName      = "can't be blank"
	errNegative    = "can't be less than 0"
	errNotPositive = "must be greater than 0"
)

type ProductForm struct {
	foodproducts.FoodProduct
	Per    float32
	Errors map[string]string
}

func (f *ProductForm) Validate() bool {
	if len(f.Name) == 0 {
		f.Errors["name"] = errNoName
	}

	if f.Calories < 0 {
		f.Errors["calories"] = errNegative
	}

	if f.Per <= 0 {
		f.Errors["per"] = errNotPositive
	}

	if f.Proteins < 0 {
		f.Errors["proteins"] = errNegative
	}

	if f.Fats < 0 {
		f.Errors["fats"] = errNegative
	}

	if f.Carbs < 0 {
		f.Errors["carbs"] = errNegative
	}

	return len(f.Errors) == 0
}

func (f *ProductForm) NormalizedProduct() foodproducts.FoodProduct {
	clone := f.FoodProduct
	clone.Normalize(f.Per)

	return clone
}
