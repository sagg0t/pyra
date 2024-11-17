package auth

import (
	"net/http"

	"pyra/internal/api/base"
)

type SignOutHandler struct {
	*base.Handler
}

func (h *SignOutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.Session(r)

	delete(session.Values, base.UserIDSessionKey)

	w.Header().Add("HX-Redirect", "/signIn")
	w.WriteHeader(http.StatusNoContent)
}
