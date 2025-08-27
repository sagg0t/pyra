package products

import (
	"database/sql"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/nutrition"
)

type EditProductHandler struct {
	*base.Handler
	productRepo nutrition.ProductRepository
}

func (h *EditProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.RequestLogger(r)

	uid, version, err := productRef(r)
	if err != nil {
		log.ErrorContext(ctx, "malformed product UID or version", "error", err)
		h.InternalServerError(w)
		return
	}

	product, err := h.productRepo.FindByRef(ctx, uid, version)
	if err != nil {
		if err == sql.ErrNoRows {
			log.WarnContext(ctx, "product not found", "uid", uid, "version", version)
			h.NotFound(w, r)
			return
		}

		log.ErrorContext(ctx, "failed to retrieve a record", "error", err)
		h.InternalServerError(w)
		return
	}

	form := ProductForm{
		Product: product,
		Per:     100,
	}

	h.Render(w, r, "edit-product", form)
}
