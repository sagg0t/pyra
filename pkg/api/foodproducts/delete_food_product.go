package foodproducts

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := api.productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.svc.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		api.log.Error("failed to delete FoodProduct", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
