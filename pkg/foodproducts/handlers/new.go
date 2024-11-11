package handlers

import (
	"net/http"

	"pyra/pkg/foodproducts"
	"pyra/pkg/foodproducts/view"
)

func (api *API) New(w http.ResponseWriter, r *http.Request) {
	form := foodproducts.ProductForm{
		FoodProduct: foodproducts.FoodProduct{},
		Per:         100,
	}

	api.Render(w, r, view.NewProduct(form))
}
