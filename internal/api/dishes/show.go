package dishes

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/internal/api/base"
	"pyra/pkg/dishes"
	"pyra/pkg/foodproducts"
)

type ShowDishHandler struct {
	*base.Handler
	svc interface {
		DishByIdFinder
		DishVersionFinder
	}

	productsRepo *foodproducts.FoodProductsRepository
}

type DishByIdFinder interface {
	FindByID(context.Context, uint64) (dishes.Dish, error)
}

type DishVersionFinder interface {
	Versions(ctx context.Context, uid string) ([]dishes.Dish, error)
}

func (h *ShowDishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	id, err := dishID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dish, err := h.svc.FindByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
	}

	versions, err := h.svc.Versions(r.Context(), dish.UID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	products, err := h.productsRepo.ForDish(r.Context(), dish.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	h.Render(w, r, "dish-details", DishDetails{dish, versions, products})
}

type DishDetails struct {
	Dish     dishes.Dish
	Versions []dishes.Dish
	Products []foodproducts.FoodProduct
}
