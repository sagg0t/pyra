// Package api - root Pyra mux.
package api

import (
	"html/template"
	"net/http"

	"pyra/internal/api/base"
	"pyra/internal/api/dishes"
	"pyra/internal/api/products"
	"pyra/pkg/db"
	"pyra/pkg/log"
)

func Mux(db db.DBTX, l *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	drivers := baseTemplate()

	baseAPI := base.NewAPI(db, drivers)
	// authAPI := auth.NewAPI(baseAPI)
	productsAPI := products.NewAPI(baseAPI)
	dishesAPI := dishes.NewAPI(baseAPI)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			if base.IsAuthenticated(r) {
				http.Redirect(w, r, "/products", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/signIn", http.StatusTemporaryRedirect)
			}
		} else {
			http.NotFound(w, r)
		}
	})

	// mux.Handle("GET /signIn", authAPI.SignIn())
	// mux.Handle("GET /auth/google", authAPI.GoogleAuthorize())
	// mux.Handle("GET /auth/google/callback", authAPI.GoogleCallback())
	// mux.Handle("GET /signOut", base.Authenticated(authAPI.SignOut()))

	mux.Handle("GET /products", base.Authenticated(productsAPI.Index()))
	mux.Handle("GET /products/{uid}/{version}", base.Authenticated(productsAPI.Show()))
	mux.Handle("GET /products/new", base.Authenticated(productsAPI.New()))
	mux.Handle("GET /products/{uid}/{version}/edit", base.Authenticated(productsAPI.Edit()))
	mux.Handle("POST /products", base.Authenticated(productsAPI.Create()))
	mux.Handle("PUT /products/{uid}/{version}", base.Authenticated(productsAPI.Update()))
	mux.Handle("DELETE /products/{uid}/{version}", base.Authenticated(productsAPI.Delete()))
	mux.Handle("POST /products/search", productsAPI.Search()) // TODO: authenticate

	mux.Handle("GET /dishes", base.Authenticated(dishesAPI.Index()))
	// mux.Handle("GET /dishes/{id}", base.Authenticated(dishesAPI.Show()))
	// mux.Handle("GET /dishes/new", base.Authenticated(dishesAPI.New()))
	// mux.Handle("POST /dishes", base.Authenticated(dishesAPI.Create()))

	return mux
}

func baseTemplate() *template.Template {
	templateDriver := template.New("drivers")

	templateDriver.Funcs(TemplateHelpers)

	template.Must(templateDriver.ParseGlob("view/layout/*.html"))
	template.Must(templateDriver.ParseGlob("view/errors/*.html"))
	template.Must(templateDriver.ParseGlob("view/components/*.html"))

	return templateDriver
}
