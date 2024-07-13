package foodproducts

import (
	"context"
	"log/slog"
	"net/http"

	fp "pyra/pkg/foodproducts"
	view "pyra/view/foodproducts"
)

func (api *API) New(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	form := fp.Form{Per: 100.0}

	if err := view.NewProduct(form).Render(ctx, w); err != nil {
		slog.Warn(err.Error())
	}
}
