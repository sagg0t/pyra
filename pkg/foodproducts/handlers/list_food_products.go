package handlers

import (
	"net/http"

	"pyra/pkg/foodproducts/view"
)

func (api *API) List(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	foodProducts, err := api.svc.Index(r.Context())
	if err != nil {
		log.Error("failed to list produces", "error", err)
		api.InternalServerError(w)
		return
	}

	api.Render(w, r, view.ProductList(foodProducts))
}
