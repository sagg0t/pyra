package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"pyra/pkg/foodproducts"
	"pyra/pkg/foodproducts/view"
)

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)
	session := api.Session(r)

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(fmt.Sprintf("failed to parse form: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form, err := paramsFromForm(r.FormValue)
	if err != nil {
		log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !form.Validate() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.NewProduct(form))
		return
	}

	newProductID, err := api.svc.Create(r.Context(), form.NormalizedProduct())
	if err != nil {
		form.Errors["base"] = err.Error()

		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.NewProduct(form))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/foodProducts/%d", newProductID), http.StatusFound)
}

func paramsFromForm(fetch func(key string) string) (foodproducts.ProductForm, error) {
	form := foodproducts.ProductForm{
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
