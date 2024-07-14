package handlers

import (
	"fmt"
	"net/http"

	"pyra/pkg/foodproducts/view"
)

func (api *API) Update(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	form, err := paramsFromForm(r.FormValue)
	if err != nil {
		log.Trace("failed to map form data", "error", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !form.Validate() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.EditProduct(form))
		return
	}

	id, err := api.productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form.ID = id

	err = api.svc.Update(r.Context(), form.NormalizedProduct())
	if err != nil {
		form.Errors["base"] = err.Error()

		w.WriteHeader(http.StatusUnprocessableEntity)
		api.Render(w, r, view.EditProduct(form))
		return
	}

	loc := fmt.Sprintf("/foodProducts/%d", id)
	w.Header().Add("HX-Redirect", loc)

	// http.Redirect(w, r, loc, http.StatusFound)
}
