package handlers

import (
	"net/http"
	"strconv"

	"pyra/pkg/foodproducts"
	"pyra/pkg/pyra"
)

type API struct {
	pyra.API

	svc foodproducts.FoodProductsRepository
}

func NewAPI(svc foodproducts.FoodProductsRepository) *API {
	return &API{
		svc: svc,
	}
}

func (api *API) productID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}
