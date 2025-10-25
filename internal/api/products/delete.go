package products

import (
	"database/sql"
	"errors"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type DeleteProductHandler struct {
	*base.Handler
	ProductRepo nutrition.ProductRepository
}

func (h *DeleteProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)

	uid, version, err := productRef(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductRepo.FindByRef(ctx, uid, version)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to find product", "error", err)
		h.InternalServerError(w)
		return
	}

	err = h.ProductRepo.Delete(ctx, product.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to delete product", "error", err)
		h.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
