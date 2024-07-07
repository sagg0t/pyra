package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	authAPI "github.com/olehvolynets/pyra/pkg/api/auth"
	foodProductsAPI "github.com/olehvolynets/pyra/pkg/api/foodproducts"
	"github.com/olehvolynets/pyra/pkg/auth"
	"github.com/olehvolynets/pyra/pkg/foodproducts"
	"github.com/olehvolynets/pyra/pkg/log"
	"github.com/olehvolynets/pyra/pkg/users"
)

func Routes(db *pgxpool.Pool, logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	authRoutes(mux, db, logger)
	foodProductsRoutes(mux, db, logger)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signIn", http.StatusTemporaryRedirect)
}

func authRoutes(mux *http.ServeMux, db *pgxpool.Pool, logger *log.Logger) {
	usersRepo := users.NewRepository(db)
	providersRepo := auth.NewProviderRepository(db)

	authSvc := auth.NewService(logger, db, providersRepo, usersRepo)

	api := authAPI.NewAPI(logger, authSvc)

	mux.HandleFunc("GET /signIn", api.SignIn)
	mux.HandleFunc("GET /auth/google", api.GoogleAuth)
	mux.HandleFunc("GET /auth/google/callback", api.GoogleCallback)
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
