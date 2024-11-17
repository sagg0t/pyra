package foodproducts

import (
	"context"
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type CreateFoodProductHandler struct {
	*base.Handler
	svc FoodProductCreator
}

type FoodProductCreator interface {
	Create(context.Context, foodproducts.FoodProduct) (id uint64, err error)
}

func (h *CreateFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)
	session := h.Session(r)

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(fmt.Sprintf("failed to parse form: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form, err := paramsFromForm(r.FormValue)
	if err != nil {
		log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !form.Validate() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-food-product", form)
		return
	}

	newProductID, err := h.svc.Create(r.Context(), form.NormalizedProduct())
	if err != nil {
		form.Errors["base"] = err.Error()

		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-food-product", form)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/foodProducts/%d", newProductID), http.StatusFound)
}
