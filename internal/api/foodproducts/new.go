package foodproducts

import (
	"net/http"

	"pyra/internal/api/base"
)

type NewFoodProductHandler struct {
	*base.Handler
}

func (h *NewFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form := ProductForm{Per: 100}

	h.Render(w, r, "new-food-product", form)
}
