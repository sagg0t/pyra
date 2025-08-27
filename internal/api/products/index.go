package products

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type ProductsHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	svc, err := nutrition.NewProductService(h.productRepo, nil)
	if err != nil {
		log.Error("failed to create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	products, err := svc.List(r.Context())
	if err != nil {
		log.Error("failed to list produces", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "product-list", products)
}
