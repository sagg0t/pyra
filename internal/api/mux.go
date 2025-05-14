package api

import (
	"html/template"
	"net/http"

	"pyra/internal/api/base"
	"pyra/internal/api/dishes"
	"pyra/internal/api/foodproducts"
	"pyra/pkg/db"
	"pyra/pkg/log"
)

func Mux(db db.DBTX, l *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	drivers := baseTemplate()

	baseAPI := base.NewAPI(db, drivers)
	// authAPI := auth.NewAPI(baseAPI)
	foodProductsAPI := foodproducts.NewAPI(baseAPI)
	dishesAPI := dishes.NewAPI(baseAPI)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			if base.IsAuthenticated(r) {
				http.Redirect(w, r, "/foodProducts", http.StatusSeeOther)
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

	mux.Handle("GET /foodProducts", base.Authenticated(foodProductsAPI.Index()))
	mux.Handle("GET /foodProducts/{id}", base.Authenticated(foodProductsAPI.Show()))
	mux.Handle("GET /foodProducts/new", base.Authenticated(foodProductsAPI.New()))
	mux.Handle("GET /foodProducts/{id}/edit", base.Authenticated(foodProductsAPI.Edit()))
	mux.Handle("POST /foodProducts", base.Authenticated(foodProductsAPI.Create()))
	mux.Handle("PUT /foodProducts/{id}", base.Authenticated(foodProductsAPI.Update()))
	mux.Handle("DELETE /foodProducts/{id}", base.Authenticated(foodProductsAPI.Delete()))
	mux.Handle("POST /foodProducts/search", foodProductsAPI.Search()) // TODO: authenticate

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
