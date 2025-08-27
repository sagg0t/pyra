package products

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"pyra/pkg/nutrition"
)

var (
	errNoName      = errors.New("can't be blank")
	errNegative    = errors.New("can't be less than 0")
	errNotPositive = errors.New("must be greater than 0")
)

func productRef(r *http.Request) (nutrition.ProductUID, nutrition.ProductVersion, error) {
	paramUID := r.PathValue("uid")
	paramVersion := r.PathValue("version")

	parsedVersion, err := strconv.ParseUint(paramVersion, 10, 64)
	if err != nil {
		return nutrition.ProductUID(""), nutrition.ProductVersion(0), err
	}

	return nutrition.ProductUID(paramUID), nutrition.ProductVersion(parsedVersion), nil
}

type ProductForm struct {
	nutrition.Product
	Per float32

	Errors productFormErrors
}

type productFormErrors struct {
	Per error
	nutrition.ProductErrors
}

func NewProductForm(fetch func(key string) string) ProductForm {
	form := ProductForm{}

	name, err := nutrition.NewProductName(fetch("name"))
	if err != nil {
		form.Errors.Name = err
	}
	form.Name = name

	calories64, err := strconv.ParseFloat(fetch("calories"), 32)
	if err != nil {
		form.Errors.Calories = err
	}

	proteins64, err := strconv.ParseFloat(fetch("proteins"), 32)
	if err != nil {
		form.Errors.Proteins = err
	}

	fats64, err := strconv.ParseFloat(fetch("fats"), 32)
	if err != nil {
		form.Errors.Fats = err
	}

	carbs64, err := strconv.ParseFloat(fetch("carbs"), 32)
	if err != nil {
		form.Errors.Carbs = err
	}

	form.Macro, form.Errors.MacroErrors = nutrition.NewMacro(
		float32(calories64),
		float32(proteins64),
		float32(fats64),
		float32(carbs64),
	)

	per64, err := strconv.ParseFloat(fetch("per"), 32)
	if err != nil {
		form.Errors.Per = err
	} else if per64 < 0 {
		form.Errors.Per = errNotPositive
	}
	form.Per = float32(per64)

	return form
}

func (e *productFormErrors) HasErrors() bool {
	perErr := e.Per != nil

	return perErr || e.ProductErrors.HasErrors()
}

func (e *productFormErrors) Error() string {
	return fmt.Errorf("per: %w\n%s", e.Per, e.ProductErrors.Error()).Error()
}

func (f *ProductForm) HasErrors() bool {
	return f.Errors.HasErrors()
}

func (f *ProductForm) NormalizedProduct() nutrition.Product {
	clone := f.Product
	clone.Normalize(f.Per)

	return clone
}
