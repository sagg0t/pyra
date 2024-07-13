package handlers

import (
	"context"
	"net/http"

	"pyra/pkg/foodproducts"
	"pyra/pkg/foodproducts/view"
)

func (api *API) New(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	ctx := context.Background()

	form := foodproducts.CreateResponse{
		CreateRequest: foodproducts.CreateRequest{Per: 100.0},
	}

	if err := view.NewProduct(form).Render(ctx, w); err != nil {
		log.Warn(err.Error())
	}
}
