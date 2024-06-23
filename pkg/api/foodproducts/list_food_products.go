package foodproducts

import (
	"fmt"
	"log/slog"
	"net/http"

	view "github.com/olehvolynets/pyra/view/foodproducts"
)

func (api *API) List(w http.ResponseWriter, r *http.Request) {
	foodProducts, err := api.svc.Index(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	component := view.ProductList(foodProducts)
	if err := component.Render(r.Context(), w); err != nil {
		slog.Warn(err.Error())
	}
}
