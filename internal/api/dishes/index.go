package dishes

import (
	"context"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/dishes"
)

type ListDishesHandler struct {
	*base.Handler
	svc DishIndexer
}

type DishIndexer interface {
	Index(context.Context) ([]dishes.Dish, error)
}

func (h *ListDishesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	dishes, err := h.svc.Index(r.Context())
	if err != nil {
		log.Error("failed to list dishes", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "dish-list", dishes)
}
