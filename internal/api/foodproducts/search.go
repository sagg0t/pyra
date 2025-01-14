package foodproducts

import (
	"context"
	"encoding/json"
	"net/http"

	"pyra/internal/api/base"
	"pyra/pkg/foodproducts"
)

type SearchFoodProductsHandler struct {
	*base.Handler
	svc FoodProductSearcher
}

type FoodProductSearcher interface {
	Search(ctx context.Context, query string) ([]foodproducts.FoodProduct, error)
}

func (h *SearchFoodProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)
	searchQuery := r.URL.Query().Get("q")

	products, err := h.svc.Search(r.Context(), searchQuery)
	if err != nil {
		log.Error("failed to fetch products", err)
		h.InternalServerError(w)
		return
	}

	searchResults := make([]searchResult, len(products))
	for idx, product := range products {
		searchResults[idx] = searchResult{
			ID:    product.ID,
			Label: product.Name,
			Value: product,
		}
	}

	res, err := json.Marshal(searchResults)
	if err != nil {
		log.Error("failed to marshal products", err)
		h.InternalServerError(w)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		log.Error("failed to write the response", err)
		h.InternalServerError(w)
	}
}

type searchResult struct {
	ID    uint64                   `json:"id"`
	Label string                   `json:"label"`
	Value foodproducts.FoodProduct `json:"value"`
}
