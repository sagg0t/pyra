package handlers

import (
	"net/http"

	"pyra/pkg/auth/view"
)

func (api *API) SignIn(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	err := view.SignIn().Render(r.Context(), w)
	if err != nil {
		log.Error("render failed", "error", err)
		api.InternalServerError(w)
	}
}
