package dishes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"pyra/internal/api/base"
	"pyra/pkg/dishes"
	"pyra/pkg/dishproducts"
	"pyra/pkg/foodproducts"
)

type NewDishHandler struct {
	*base.Handler
	svc         DishCreator
	productsSvc ProductAllFinder
}

type DishCreator interface {
	Create(context.Context, dishes.Dish) (id uint64, err error)
}

type ProductAllFinder interface {
	FindAllByIds(ctx context.Context, ids []uint64) ([]foodproducts.FoodProduct, error)
}

func (h *NewDishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := h.RequestLogger(r)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error parsing params")
		return
	}

	productIDstr := r.PostForm["product"]
	productAmounts := r.PostForm["amounts"]

	if len(productIDstr) != len(productAmounts) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Number of product IDs doesn't match the number of amounts")
		return
	}

	ingredients := make([]dishproducts.DishProduct, len(productIDstr))
	productIDs := make([]uint64, len(productIDstr))

	for i := 0; i < len(productIDstr); i++ {
		id, err := strconv.ParseUint(productIDstr[i], 10, 64)
		if err != nil {
		}

		amount, err := strconv.ParseFloat(productIDstr[i], 64)
		if err != nil {
		}

		ingredients[i].ProductID = id
		ingredients[i].Amount = float32(amount)

		productIDs[i] = id
	}

	products, err := h.productsSvc.FindAllByIds(r.Context(), productIDs)
	if len(products) != len(productIDs) {
		l.Debug("some of the requested products weren't found")

		form := DishForm{}
		w.WriteHeader(http.StatusNotFound)
		h.Render(w, r, "new-dish", form)
		return
	}

	dish := dishes.Dish{Name: r.PostForm.Get("name")}
	dish.SetIngredients(ingredients)

	for i := 0; i < len(productIDs); i++ {
		productId := productIDs[i]
		amount := productAmounts[i]

		_ = productId
		_ = amount
	}

	id, err := h.svc.Create(r.Context(), dish)
	if err != nil {
		form := DishForm{}
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-dish", form)
	}

	http.Redirect(w, r, fmt.Sprintf("/dishes/%d", id), http.StatusFound)
}
