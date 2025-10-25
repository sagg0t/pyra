package products

import (
	"errors"
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type CreateProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *CreateProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)
	session := h.Session(r)

	err := r.ParseForm()
	if err != nil {
		session.AddFlash(fmt.Sprintf("failed to parse form: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := NewProductForm(r.FormValue)
	product := form.BuildProduct()

	if form.HasErrors() {
		log.DebugContext(ctx, "create product validation error", "error", form.Errors)
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-product", form)
		return
	}

	err = nutrition.CreateProduct(ctx, h.productRepo, &product)
	if err != nil && !errors.Is(err, nutrition.ErrProductInvalid) {
		log.DebugContext(ctx, "failed to save product", "error", err)
		h.InternalServerError(w)
		return
	}

	if product.HasErrors() {
		form.SetProductErrors(product.Errors)
		session.AddFlash("failed to create a product")
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-product", form)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf("/products/%s/%d", product.UID, product.Version),
		http.StatusFound)
}
