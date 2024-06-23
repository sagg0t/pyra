package foodproducts

import fp "github.com/olehvolynets/pyra/pkg/foodproducts"

type API struct {
	svc fp.FoodProductsService
}

func NewAPI(svc fp.FoodProductsService) *API {
	return &API{
		svc: svc,
	}
}
