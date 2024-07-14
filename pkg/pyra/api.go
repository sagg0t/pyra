package pyra

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"

	"pyra/pkg/auth"
	"pyra/pkg/log"
	"pyra/pkg/session"
	"pyra/pkg/users"
)

var ErrNoUsesr = errors.New("no current user")

type API struct {
	UserSvc users.UserRepository
}

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

type RenderContext struct {
	User     *users.User
	UserPath []string
}
