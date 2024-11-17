package foodproducts

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type FoodProductHanler struct {
	*base.Handler
	svc FoodProductByIdFinder
}

type FoodProductByIdFinder interface {
	FindById(context.Context, uint64) (foodproducts.FoodProduct, error)
}

func (h *FoodProductHanler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.svc.FindById(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.Error("failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "food-product-details", product)
}
