package handlers

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"

	"pyra/pkg/dishes/view"
	"pyra/pkg/log"
)

func (api *API) Details(w http.ResponseWriter, r *http.Request) {
	log := log.FromContext(r.Context())

	id, err := api.dishID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dish, err := api.svc.FindByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
	}

	versions, err := api.svc.Versions(r.Context(), dish.UID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		api.InternalServerError(w)
		return
	}

	products, err := api.productsRepo.ForDish(r.Context(), dish.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		log.Error("failed to retrieve a record", "error", err)
		api.InternalServerError(w)
		return
	}

	api.Render(w, r, view.DishDetails(dish, versions, products))
}
