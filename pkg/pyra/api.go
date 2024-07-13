package pyra

import (
	"net/http"

	"github.com/a-h/templ"

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

func (api *API) NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (api *API) InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (api *API) Render(w http.ResponseWriter, r *http.Request, comp templ.Component) {
	ctx := r.Context()

	if err := comp.Render(ctx, w); err != nil {
		logger := api.RequestLogger(r)
		logger.ErrorContext(ctx, "render failed", "error", err)
		api.InternalServerError(w)
	}
}
