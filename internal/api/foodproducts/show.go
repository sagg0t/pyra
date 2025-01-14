package foodproducts

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/internal/api/base"
	"pyra/pkg/dishes"
	"pyra/pkg/foodproducts"
)

type FoodProductHandler struct {
	*base.Handler
	svc interface {
		FoodProductByIdFinder
		FoodProductVersionFinder
	}

	dishSvc interface {
		FindAllByProductID(ctx context.Context, productID uint64) ([]dishes.Dish, error)
	}
}

type FoodProductByIdFinder interface {
	FindById(context.Context, uint64) (foodproducts.FoodProduct, error)
}

type FoodProductVersionFinder interface {
	Versions(context.Context, string) ([]foodproducts.FoodProduct, error)
}

func (h *FoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	versions, err := h.svc.Versions(r.Context(), product.UID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	usedInDishes, err := h.dishSvc.FindAllByProductID(r.Context(), product.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	details := FoodProductDetails{
		Product:      product,
		Versions:     versions,
		UsedInDishes: usedInDishes,
	}

	h.Render(w, r, "food-product-details", details)
}

type FoodProductDetails struct {
	Product      foodproducts.FoodProduct
	Versions     []foodproducts.FoodProduct
	UsedInDishes []dishes.Dish
}
