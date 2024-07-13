package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	authAPI "pyra/pkg/api/auth"
	foodProductsAPI "pyra/pkg/api/foodproducts"
	"pyra/pkg/auth"
	"pyra/pkg/foodproducts"
	"pyra/pkg/log"
	"pyra/pkg/users"
)

func Routes(db *pgxpool.Pool, logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	authRoutes(mux, db, logger)
	foodProductsRoutes(mux, db, logger)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/foodProducts", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/signIn", http.StatusTemporaryRedirect)
	}
}

func authRoutes(mux *http.ServeMux, db *pgxpool.Pool, logger *log.Logger) {
	usersRepo := users.NewRepository(db)
	providersRepo := auth.NewProviderRepository(db)

	authSvc := auth.NewService(logger, db, providersRepo, usersRepo)

	api := authAPI.NewAPI(logger, authSvc)

	mux.Handle("GET /signIn", auth.NotAuthenticated(api.SignIn))
	mux.Handle("GET /auth/google", auth.NotAuthenticated(api.GoogleAuth))
	mux.Handle("GET /auth/google/callback", auth.NotAuthenticated(api.GoogleCallback))
}

func foodProductsRoutes(mux *http.ServeMux, db *pgxpool.Pool, logger *log.Logger) {
	service := foodproducts.NewDB(db)
	api := foodProductsAPI.NewAPI(logger, service)

	mux.Handle("GET /foodProducts", auth.Authenticated(api.List))
	mux.Handle("GET /foodProducts/{id}", auth.Authenticated(api.Show))
	mux.Handle("GET /foodProducts/new", auth.Authenticated(api.New))
	mux.Handle("POST /foodProducts", auth.Authenticated(api.Create))
	mux.Handle("DELETE /foodProducts/{id}", auth.Authenticated(api.Delete))
}
