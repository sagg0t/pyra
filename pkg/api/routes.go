package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	foodProductsAPI "github.com/olehvolynets/pyra/pkg/api/foodproducts"
	"github.com/olehvolynets/pyra/pkg/foodproducts"
)

func Routes(db *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	foodProductsRoutes(mux, db)

	return mux
}

func foodProductsRoutes(mux *http.ServeMux, db *pgxpool.Pool) {
	service := foodproducts.NewService(db)
	api := foodProductsAPI.NewAPI(service)

	mux.HandleFunc("GET /foodProducts", api.List)
	mux.HandleFunc("GET /foodProducts/{id}", api.Show)
	mux.HandleFunc("GET /foodProducts/new", api.New)
	mux.HandleFunc("POST /foodProducts", api.Create)
}
