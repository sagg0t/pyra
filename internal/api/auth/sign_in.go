package auth

import (
	"net/http"

	"pyra/internal/api/base"
)

type SignInHandler struct {
	*base.Handler
}

func (h *SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Render(w, r, "sign-in", nil)
}
