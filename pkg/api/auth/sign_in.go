package auth

import (
	"net/http"

	view "pyra/view/auth"
)

func (api *API) SignIn(w http.ResponseWriter, r *http.Request) {
	err := view.SignIn().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
