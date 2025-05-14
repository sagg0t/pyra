package foodproducts

import (
	"database/sql"
	"errors"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type FoodProductHandler struct {
	*base.Handler

	productRepo nutrition.ProductRepository
	dishRepo    nutrition.DishRepository
}

// GET /foodProducts/:id
func (h *FoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)

	id, err := productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	versions, err := h.productRepo.Versions(ctx, product.UID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.ErrorContext(ctx, "failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	usedInDishes, err := h.dishRepo.FindAllByProductID(ctx, product.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.ErrorContext(ctx, "failed to retrieve a record", "error", err)
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
	Product      nutrition.Product
	Versions     []nutrition.Product
	UsedInDishes []nutrition.Dish
}
