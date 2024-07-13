package foodproducts

import (
	"fmt"
	"net/http"
	"strconv"

	"pyra/pkg/foodproducts"
)

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		api.log.Trace("failed to parse form", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	reqData, err := paramsFromForm(r.FormValue)
	if err != nil {
		api.log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	validator := foodproducts.NewCreateValidator(reqData)
	validator.Validate()
	if err = validator.Err(); err != nil {
		api.log.Trace("validation error", "error", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, err)
		return
	}

	newProductID, err := api.svc.Create(r.Context(), reqData)
	if err != nil {
		api.log.Trace("validation error", "error", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf("/foodProducts/%d", newProductID),
		http.StatusMovedPermanently)
}

func paramsFromForm(fetch func(key string) string) (foodproducts.Form, error) {
	reqData := foodproducts.Form{
		Name: fetch("name"),
	}

	calories64, err := strconv.ParseFloat(fetch("calories"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Calories = float32(calories64)

	per64, err := strconv.ParseUint(fetch("per"), 10, 32)
	if err != nil {
		return reqData, err
	}
	reqData.Per = uint32(per64)

	proteins64, err := strconv.ParseFloat(fetch("proteins"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Proteins = float32(proteins64)

	fats64, err := strconv.ParseFloat(fetch("fats"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Fats = float32(fats64)

	carbs64, err := strconv.ParseFloat(fetch("carbs"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Carbs = float32(carbs64)

	return reqData, nil
}
