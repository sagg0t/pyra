package foodproducts

import (
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type FoodProductsHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *FoodProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	svc, err := nutrition.NewProductService(h.productRepo)
	if err != nil {
		log.Error("failed to create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	foodProducts, err := svc.List(r.Context())
	if err != nil {
		log.Error("failed to list produces", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "food-product-list", foodProducts)
}
