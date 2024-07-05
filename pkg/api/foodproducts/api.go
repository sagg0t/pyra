package foodproducts

import (
	"net/http"
	"strconv"

	fp "github.com/olehvolynets/pyra/pkg/foodproducts"
	"github.com/olehvolynets/pyra/pkg/log"
)

type API struct {
	svc fp.FoodProductsDB
	log *log.Logger
}

func NewAPI(logger *log.Logger, svc fp.FoodProductsDB) *API {
	return &API{
		svc: svc,
		log: logger,
	}
}

func (api *API) productID(r *http.Request) (uint64, error) {
	paramID := r.PathValue("id")
	return strconv.ParseUint(paramID, 10, 64)
}
