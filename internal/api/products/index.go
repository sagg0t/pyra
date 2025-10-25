package products

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type ProductsHandler struct {
	*base.Handler
	ProductRepo nutrition.ProductRepository
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	products, err := nutrition.ListProducts(r.Context(), h.ProductRepo)
	if err != nil {
		log.Error("failed to list produces", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "product-list", products)
}
