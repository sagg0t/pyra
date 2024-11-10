package handlers

import (
	"net/http"

	"pyra/pkg/dishes"
	"pyra/pkg/dishes/view"
)

func (api *API) NewDish(w http.ResponseWriter, r *http.Request) {
	form := dishes.DishForm{}
	api.Render(w, r, view.NewDish(form))
}
