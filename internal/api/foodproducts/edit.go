package foodproducts

import (
	"database/sql"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type EditFoodProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *EditFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		log.ErrorContext(ctx, "failed to extract ID from URI", "error", err)
		h.InternalServerError(w)
		return
	}

	product, err := h.productRepo.FindByID(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.WarnContext(ctx, "food product not found", "id", id)
			h.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	form := ProductForm{
		Product: product,
		Per:     100,
	}

	h.Render(w, r, "edit-food-product", form)
}
