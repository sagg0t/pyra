package handlers

import (
	"net/http"

	"pyra/pkg/auth/view"
)

func (api *API) SignIn(w http.ResponseWriter, r *http.Request) {
	api.Render(w, r, view.SignIn())
}
