package foodproducts

import (
	"database/sql"
	"errors"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type DeleteFoodProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *DeleteFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to delete FoodProduct", "error", err)
		h.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
