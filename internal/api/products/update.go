package products

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type UpdateProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
	dishRepo    nutrition.DishRepository
}

// PUT /products/{id}
func (h *UpdateProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)
	session := h.Session(r)

	form := NewProductForm(r.FormValue)
	if form.HasErrors() {
		log.TraceContext(ctx, "update product form error", "error", form.Errors.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-product", form)
		return
	}

	uid, version, err := productRef(r)
	if err != nil {
		log.TraceContext(ctx, "malformed product UID or version", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	svc, err := nutrition.NewProductService(h.productRepo, h.dishRepo)
	if err != nil {
		log.TraceContext(ctx, "couldn't create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	updateInfo := nutrition.UpdateProductInfo{
		UID:     uid,
		Version: version,
		Name:    form.Name,
		Macro:   form.NormalizedProduct().Macro,
	}

	product, err := svc.Update(r.Context(), updateInfo)
	if err != nil {
		errMsg := fmt.Sprintf("couldn't update the product: %s", err.Error())
		log.DebugContext(ctx, errMsg)

		if errors.Is(err, sql.ErrNoRows) {
			h.NotFound(w, r)
			return
		}

		session.AddFlash(errMsg)
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "edit-product", form)
		return
	}

	loc := fmt.Sprintf("/products/%s/%d", product.UID, product.Version)
	w.Header().Add("HX-Redirect", loc)

	// http.Redirect(w, r, loc, http.StatusFound)
}
