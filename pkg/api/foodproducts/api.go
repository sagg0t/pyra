package foodproducts

import (
	"net/http"
	"strconv"

	fp "pyra/pkg/foodproducts"
	"pyra/pkg/log"
)

type API struct {
	svc fp.FoodProductsRepository
	log *log.Logger
}

func NewAPI(logger *log.Logger, svc fp.FoodProductsRepository) *API {
	return &API{
		svc: svc,
		log: logger,
	}
}

func (api *API) productID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}
