package foodproduct

import (
	"context"
	"net/http"

	view "github.com/olehvolynets/pyra/internal/view/food_product"
)

func ServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		component := view.ProductList()
		component.Render(context.TODO(), w)
	})

	return mux
}
