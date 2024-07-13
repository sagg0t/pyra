package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"pyra/pkg/auth"
	authAPI "pyra/pkg/auth/handlers"
	"pyra/pkg/foodproducts"
	foodProductsAPI "pyra/pkg/foodproducts/handlers"
	"pyra/pkg/log"
	"pyra/pkg/users"
)

func Routes(db *pgxpool.Pool, logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	authRoutes(mux, db, logger)
	foodProductsRoutes(mux, db)

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

	api := authAPI.NewAPI(authSvc)

	mux.Handle("GET /signIn", auth.NotAuthenticated(api.SignIn))
	mux.Handle("GET /auth/google", auth.NotAuthenticated(api.GoogleAuth))
	mux.Handle("GET /auth/google/callback", auth.NotAuthenticated(api.GoogleCallback))
}

func foodProductsRoutes(mux *http.ServeMux, db *pgxpool.Pool) {
	service := foodproducts.NewDB(db)
	api := foodProductsAPI.NewAPI(service)

	mux.Handle("GET /foodProducts", auth.Authenticated(api.List))
	mux.Handle("GET /foodProducts/{id}", auth.Authenticated(api.Show))
	mux.Handle("GET /foodProducts/new", auth.Authenticated(api.New))
	mux.Handle("GET /foodProducts/{id}/edit", auth.Authenticated(api.Edit))
	mux.Handle("POST /foodProducts", auth.Authenticated(api.Create))
	mux.Handle("DELETE /foodProducts/{id}", auth.Authenticated(api.Delete))
}
