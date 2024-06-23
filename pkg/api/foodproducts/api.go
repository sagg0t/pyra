package foodproducts

import (
	fp "github.com/olehvolynets/pyra/pkg/foodproducts"
	"github.com/olehvolynets/pyra/pkg/log"
)

type API struct {
	svc fp.FoodProductsService
	log *log.Logger
}

func NewAPI(logger *log.Logger, svc fp.FoodProductsService) *API {
	return &API{
		svc: svc,
		log: logger,
	}
}
