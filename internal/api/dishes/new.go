package dishes

import (
	"net/http"

	"pyra/internal/api/base"
)

type NewDishHandler struct {
	*base.Handler
}

func (h *NewDishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form := DishForm{}
	h.Render(w, r, "new-dish", form)
}
