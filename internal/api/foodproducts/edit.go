package foodproducts

import (
	"net/http"

	"pyra/internal/api/base"
)

type EditFoodProductHandler struct {
	*base.Handler
	svc FoodProductByIdFinder
}

func (h *EditFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		log.ErrorContext(r.Context(), "failed to extract ID from URI")
		h.InternalServerError(w)
		return
	}

	product, err := h.svc.FindById(r.Context(), id)
	if err != nil {
		log.WarnContext(r.Context(), "food product not found", "id", id)
		h.NotFound(w, r)
		return
	}

	form := ProductForm{
		FoodProduct: product,
		Per:         100,
	}

	h.Render(w, r, "edit-food-product", form)
}
