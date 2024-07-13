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

	reqData, err := paramsFromForm(r.FormValue)
	if err != nil {
		log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	validator := foodproducts.NewCreateValidator(reqData)
	validator.Validate()
	validationErrors := validator.Err()
	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.NewProduct(foodproducts.CreateResponse{
			CreateRequest: reqData,
			Errors:        validationErrors,
		}))
		return
	}

	newProductID, err := api.svc.Create(r.Context(), reqData)
	if err != nil {
		validationErrors["base"] = err.Error()

		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.NewProduct(foodproducts.CreateResponse{
			CreateRequest: reqData,
			Errors:        validationErrors,
		}))
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/foodProducts/%d", newProductID), http.StatusFound)
}

func paramsFromForm(fetch func(key string) string) (foodproducts.CreateRequest, error) {
	reqData := foodproducts.CreateRequest{
		Name: fetch("name"),
	}

	calories64, err := strconv.ParseFloat(fetch("calories"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Calories = float32(calories64)

	per64, err := strconv.ParseFloat(fetch("per"), 32)
	if err != nil {
		return reqData, err
	}
	reqData.Per = float32(per64)

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
