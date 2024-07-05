package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	foodProductsAPI "github.com/olehvolynets/pyra/pkg/api/foodproducts"
	"github.com/olehvolynets/pyra/pkg/foodproducts"
	"github.com/olehvolynets/pyra/pkg/log"
)

func Routes(db *pgxpool.Pool, logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	foodProductsRoutes(mux, db, logger)

	return mux
}

func foodProductsRoutes(mux *http.ServeMux, db *pgxpool.Pool, logger *log.Logger) {
	service := foodproducts.NewDB(db)
	api := foodProductsAPI.NewAPI(logger, service)

	mux.HandleFunc("GET /foodProducts", api.List)
	mux.HandleFunc("GET /foodProducts/{id}", api.Show)
	mux.HandleFunc("GET /foodProducts/new", api.New)
	mux.HandleFunc("POST /foodProducts", api.Create)
	mux.HandleFunc("DELETE /foodProducts/{id}", api.Delete)
}
