package pyra

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/a-h/templ"

	"pyra/pkg/auth"
	"pyra/pkg/log"
	"pyra/pkg/session"
	"pyra/pkg/users"
)

var ErrNoUsesr = errors.New("no current user")

type API struct {
	templateDrivers *template.Template
	UserSvc         users.UserRepository
}

func NewAPI(t *template.Template) *API {
	return &API{
		templateDrivers: t,
	}
}

// INFO: done
func (api *API) RequestLogger(r *http.Request) *log.Logger {
	return log.FromContext(r.Context())
}

// INFO: done
func (api *API) Session(r *http.Request) *session.Session {
	return session.FromContext(r.Context())
}

// INFO: done
func (api *API) NotFound(w http.ResponseWriter, r *http.Request) {
	api.templateDrivers.ExecuteTemplate(w, "not-found-error", nil)
}

// INFO: done
func (api *API) InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// INFO: done
func (api *API) Render(w http.ResponseWriter, r *http.Request, comp templ.Component) {
	ctx := r.Context()

	if err := comp.Render(ctx, w); err != nil {
		logger := api.RequestLogger(r)
		logger.ErrorContext(ctx, "render failed", "error", err)
		api.InternalServerError(w)
	}
}

// INFO: done
func (api *API) CurrentUser(r *http.Request) (*users.User, error) {
	s := api.Session(r)
	userId, ok := s.Values[auth.UserIDSessionKey]
	if !ok {
		return nil, ErrNoUsesr
	}

	user, err := api.UserSvc.FindById(r.Context(), userId.(uint64))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
