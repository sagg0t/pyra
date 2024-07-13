package foodproducts

import (
	"net/http"

	"pyra/pkg/log"
	view "pyra/view/foodproducts"
)

func (api *API) List(w http.ResponseWriter, r *http.Request) {
	l := log.FromContext(r.Context())

	foodProducts, err := api.svc.Index(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		l.Error("failed to list produces", "error", err)
		return
	}

	component := view.ProductList(foodProducts)
	if err := component.Render(r.Context(), w); err != nil {
		l.Warn(err.Error())
	}
}
