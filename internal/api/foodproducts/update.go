package foodproducts

import (
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type UpdateFoodProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

// PUT /foodProducts/{id}
func (h *UpdateFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)
	session := h.Session(r)

	form := NewProductForm(r.FormValue)
	if form.HasErrors() {
		log.TraceContext(ctx, form.Errors.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-food-product", form)
		return
	}

	id, err := productID(r)
	if err != nil {
		log.TraceContext(ctx, "malformed product ID", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	svc, err := nutrition.NewProductService(h.productRepo)
	if err != nil {
		log.TraceContext(ctx, "couldn't create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	updateInfo := nutrition.UpdateProductInfo{
		ID:    id,
		Name:  form.Name,
		Macro: form.NormalizedProduct().Macro,
	}

	product, err := svc.Update(r.Context(), updateInfo)
	if err != nil {
		session.AddFlash(fmt.Sprintf("couldn't update the product: %s", err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-food-product", form)
		return
	}

	loc := fmt.Sprintf("/foodProducts/%d", product.ID)
	w.Header().Add("HX-Redirect", loc)

	// http.Redirect(w, r, loc, http.StatusFound)
}
