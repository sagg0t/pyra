package handlers

import (
	"encoding/json"
	"net/http"
)

func (api *API) Search(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)
	searchQuery := r.URL.Query().Get("q")

	products, err := api.svc.Search(r.Context(), searchQuery)
	if err != nil {
		log.Error("failed to fetch products", err)
		api.InternalServerError(w)
	}

	res, err := json.Marshal(products)
	if err != nil {
		log.Error("failed to marshal products", err)
		api.InternalServerError(w)
	}

	_, err = w.Write(res)
	if err != nil {
		log.Error("failed to write the response", err)
		api.InternalServerError(w)
	}
}
