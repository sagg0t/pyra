package foodproduct

import (
	"net/http"
)

func ServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", listFoodProducts)
	mux.HandleFunc("POST /", createFoodProduct)

	return mux
}
