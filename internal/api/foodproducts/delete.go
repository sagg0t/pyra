package foodproducts

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/internal/api/base"
)

type DeleteFoodProductHandler struct {
	*base.Handler
	svc FoodProductDeleter
}

type FoodProductDeleter interface {
	Delete(ctx context.Context, id uint64) error
}

func (h *DeleteFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.svc.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.Error("failed to delete FoodProduct", "error", err)
		h.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
