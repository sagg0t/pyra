package foodproducts

import (
	"context"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type FoodProductsHandler struct {
	*base.Handler
	svc FoodProductIndexer
}

type FoodProductIndexer interface {
	Index(context.Context) ([]foodproducts.FoodProduct, error)
}

func (h *FoodProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	foodProducts, err := h.svc.Index(r.Context())
	if err != nil {
		log.Error("failed to list produces", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "food-product-list", foodProducts)
}
