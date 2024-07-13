package handlers

import (
	"net/http"

	"pyra/pkg/foodproducts"
	"pyra/pkg/foodproducts/view"
)

func (api *API) Edit(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	id, err := api.productID(r)
	if err != nil {
		log.ErrorContext(r.Context(), "failed to extract ID from URI")
		api.InternalServerError(w)
		return
	}

	product, err := api.svc.FindById(r.Context(), id)
	if err != nil {
		log.WarnContext(r.Context(), "food product not found", "id", id)
		api.NotFound(w, r)
		return
	}

	form := foodproducts.CreateResponse{
		CreateRequest: foodproducts.CreateRequest{
			Name:     product.Name,
			Calories: product.Calories,
			Per:      100.0,
			Proteins: product.Proteins,
			Fats:     product.Fats,
			Carbs:    product.Carbs,
		},
	}

	api.Render(w, r, view.EditProduct(form))
}
