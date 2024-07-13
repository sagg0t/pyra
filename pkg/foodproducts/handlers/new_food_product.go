package handlers

import (
	"net/http"

	"pyra/pkg/foodproducts"
	"pyra/pkg/foodproducts/view"
)

func (api *API) New(w http.ResponseWriter, r *http.Request) {
	form := foodproducts.CreateResponse{
		CreateRequest: foodproducts.CreateRequest{Per: 100.0},
	}

	api.Render(w, r, view.NewProduct(form))
}
