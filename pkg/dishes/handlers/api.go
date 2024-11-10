package handlers

import (
	"net/http"
	"strconv"

	"pyra/pkg/dishes"
	"pyra/pkg/foodproducts"
	"pyra/pkg/pyra"
)

type API struct {
	pyra.API

	svc          dishes.Repository
	productsRepo foodproducts.FoodProductsRepository
}

func NewAPI(base pyra.API, repo dishes.Repository, productsRepo foodproducts.FoodProductsRepository) *API {
	return &API{
		API:          base,
		svc:          repo,
		productsRepo: productsRepo,
	}
}

func (api *API) dishID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}
