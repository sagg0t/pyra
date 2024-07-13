package foodproducts

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	view "pyra/view/foodproducts"
)

func (api *API) Show(w http.ResponseWriter, r *http.Request) {
	id, err := api.productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := api.svc.FindById(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		api.log.Error("failed to retrieve a record", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := view.ProductDetails(product).Render(r.Context(), w); err != nil {
		api.log.Warn(err.Error())
	}
}
