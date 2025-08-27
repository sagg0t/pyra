package products

import (
	"fmt"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type CreateProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

// POST /products
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
	if form.HasErrors() {
		log.DebugContext(ctx, "create product validation error", "error", form.Errors)
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-product", form)
		return
	}

	svc, err := nutrition.NewProductService(h.productRepo, nil)
	if err != nil {
		log.ErrorContext(ctx, "failed to create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	createInfo := nutrition.CreateProductInfo{
		Name:  form.Name,
		Macro: form.NormalizedProduct().Macro,
	}

	product, err := svc.Create(ctx, createInfo)
	if err != nil {
		log.DebugContext(ctx, "failed to save product", "error", err)
		session.AddFlash(fmt.Sprintf("failed to create a product: %s", err.Error()))

		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-product", form)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf("/products/%s/%d", product.UID, product.Version),
		http.StatusFound)
}
