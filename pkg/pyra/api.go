package pyra

import (
	"net/http"

	"pyra/pkg/log"
	"pyra/pkg/session"
)

type API struct{}

func (api *API) RequestLogger(r *http.Request) *log.Logger {
	return log.FromContext(r.Context())
}

func (api *API) Session(r *http.Request) *session.Session {
	return session.FromContext(r.Context())
}

func (api *API) InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
