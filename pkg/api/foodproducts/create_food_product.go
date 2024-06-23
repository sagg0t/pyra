package foodproducts

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (api *API) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		api.log.Error("failed to parse form", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	reqData, err := paramsFromForm(r.FormValue)
	if err != nil {
		api.log.Error("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	bytes, _ := json.MarshalIndent(reqData, "", "	")
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

type CreateRequest struct {
	Name string

	Calories float32
	Per      uint32

	Proteins float32
	Fats     float32
	Carbs    float32
}

func paramsFromForm(fetch func(key string) string) (CreateRequest, error) {
	reqData := CreateRequest{
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
