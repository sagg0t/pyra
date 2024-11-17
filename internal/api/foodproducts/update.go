package foodproducts

import (
	"context"
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type UpdateFoodProductHandler struct {
	*base.Handler
	svc FoodProductUpdater
}

type FoodProductUpdater interface {
	Update(context.Context, foodproducts.FoodProduct) error
}

func (h *UpdateFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	form, err := paramsFromForm(r.FormValue)
	if err != nil {
		log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !form.Validate() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-food-product", form)
		return
	}

	id, err := productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form.ID = id

	err = h.svc.Update(r.Context(), form.NormalizedProduct())
	if err != nil {
		form.Errors["base"] = err.Error()

		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-food-product", form)
		return
	}

	loc := fmt.Sprintf("/foodProducts/%d", id)
	w.Header().Add("HX-Redirect", loc)

	// http.Redirect(w, r, loc, http.StatusFound)
}
