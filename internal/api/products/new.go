package products

import (
	"net/http"

	"pyra/internal/api/base"
)

type NewProductHandler struct {
	*base.Handler
}

func (h *NewProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form := ProductForm{Per: 100}

	h.Render(w, r, "new-product", form)
}
