package handlers

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/pkg/foodproducts/view"
)

func (api *API) Show(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	id, err := api.productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := api.svc.FindById(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.Error("failed to retrieve a record", "error", err)
		api.InternalServerError(w)
		return
	}

	if err := view.ProductDetails(product).Render(r.Context(), w); err != nil {
		log.Warn(err.Error())
	}
}
