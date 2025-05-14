package foodproducts

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type CreateFoodProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

// POST /foodProducts
func (h *CreateFoodProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-food-product", form)
		return
	}

	svc, err := nutrition.NewProductService(h.productRepo)
	if err != nil {
		log.ErrorContext(ctx, "failed to create ProductService", "error", err)
		h.InternalServerError(w)
		return
	}

	createInfo := nutrition.CreateProductInfo{
		UID:   nutrition.ProductUID(uuid.New().String()),
		Name:  form.Name,
		Macro: form.NormalizedProduct().Macro,
	}

	product, err := svc.Create(r.Context(), createInfo)
	if err != nil {
		session.AddFlash(fmt.Sprintf("failed to create a product: %s", err.Error()))

		w.WriteHeader(http.StatusUnprocessableEntity)
		h.Render(w, r, "new-food-product", form)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf("/foodProducts/%d", product.ID),
		http.StatusFound)
}
