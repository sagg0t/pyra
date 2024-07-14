package handlers

import (
	"net/http"

	"pyra/pkg/auth"
)

func (api *API) SignOut(w http.ResponseWriter, r *http.Request) {
	session := api.Session(r)

	delete(session.Values, auth.UserIDSessionKey)

	w.Header().Add("HX-Redirect", "/signIn")
	w.WriteHeader(http.StatusNoContent)
}
