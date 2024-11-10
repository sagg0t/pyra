package handlers

import (
	"net/http"

	"pyra/pkg/dishes/view"
)

func (api *API) List(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	dishes, err := api.svc.Index(r.Context())
	if err != nil {
		log.Error("failed to list dishes", "error", err)
		api.InternalServerError(w)
		return
	}

	api.Render(w, r, view.DishList(dishes))
}
