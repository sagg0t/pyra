package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"pyra/pkg/auth"
	authAPI "pyra/pkg/auth/handlers"
	"pyra/pkg/dishes"
	dishAPI "pyra/pkg/dishes/handlers"
	"pyra/pkg/foodproducts"
	foodProductsAPI "pyra/pkg/foodproducts/handlers"
	"pyra/pkg/log"
	"pyra/pkg/pyra"
	"pyra/pkg/users"
)

func Routes(db *pgxpool.Pool, logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	usersRepo := users.NewRepository(db)
	providersRepo := auth.NewProviderRepository(db)
	foodProductRepo := foodproducts.NewRepository(db)
	dishesRepo := dishes.NewRepository(db)

	authSvc := auth.NewService(logger, db, providersRepo, usersRepo)

	baseHandler := pyra.API{
		UserSvc: usersRepo,
	}
	authHandler := authAPI.NewAPI(authSvc)
	foodProductsHandler := foodProductsAPI.NewAPI(baseHandler, foodProductRepo)
	dishesHandler := dishAPI.NewAPI(baseHandler, dishesRepo, foodProductRepo)

	mux.HandleFunc("/", rootHandler)

	mux.Handle("GET /signIn", auth.NotAuthenticated(authHandler.SignIn))
	mux.Handle("GET /auth/google", auth.NotAuthenticated(authHandler.GoogleAuth))
	mux.Handle("GET /auth/google/callback", auth.NotAuthenticated(authHandler.GoogleCallback))
	mux.Handle("POST /signOut", auth.Authenticated(authHandler.SignOut))

	mux.Handle("GET /foodProducts", auth.Authenticated(foodProductsHandler.List))
	mux.Handle("GET /foodProducts/{id}", auth.Authenticated(foodProductsHandler.Show))
	mux.Handle("GET /foodProducts/new", auth.Authenticated(foodProductsHandler.New))
	mux.Handle("GET /foodProducts/{id}/edit", auth.Authenticated(foodProductsHandler.Edit))
	mux.Handle("POST /foodProducts", auth.Authenticated(foodProductsHandler.Create))
	mux.Handle("PUT /foodProducts/{id}", auth.Authenticated(foodProductsHandler.Update))
	mux.Handle("DELETE /foodProducts/{id}", auth.Authenticated(foodProductsHandler.Delete))
	mux.Handle("POST /foodProducts/search", http.HandlerFunc(foodProductsHandler.Search))

	mux.Handle("GET /dishes", auth.Authenticated(dishesHandler.List))
	mux.Handle("GET /dishes/{id}", auth.Authenticated(dishesHandler.Details))
	mux.Handle("GET /dishes/new", auth.Authenticated(dishesHandler.NewDish))

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if auth.IsAuthenticated(r) {
			http.Redirect(w, r, "/foodProducts", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signIn", http.StatusTemporaryRedirect)
		}
	} else {
		http.NotFound(w, r)
	}
}
