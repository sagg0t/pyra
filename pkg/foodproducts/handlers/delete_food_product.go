package handlers

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	id, err := api.productID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.svc.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}

		log.Error("failed to delete FoodProduct", "error", err)
		api.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
