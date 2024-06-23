package foodproducts

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/olehvolynets/pyra/view/foodproducts"
)

func (api *API) New(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if err := foodproducts.NewProduct().Render(ctx, w); err != nil {
		slog.Warn(err.Error())
	}
}
