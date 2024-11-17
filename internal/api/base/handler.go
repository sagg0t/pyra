package base

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"pyra/pkg/log"
	"pyra/pkg/session"
	"pyra/pkg/users"
)

var ErrNoUsesr = errors.New("no current user")

type Handler struct {
	template *template.Template
	userSvc  UserFinder
}

type UserFinder interface {
	FindById(context.Context, uint64) (users.User, error)
}

func (h *Handler) Render(w http.ResponseWriter, r *http.Request, name string, data any) {
	if err := h.template.ExecuteTemplate(w, name, data); err != nil {
		panic(fmt.Errorf("render failed: %w", err))
	}
}

func (h *Handler) ExpandTemplate(paths ...string) error {
	_, err := h.template.ParseFiles(paths...)

	return err
}

func (h *Handler) Session(r *http.Request) *session.Session {
	return session.FromContext(r.Context())
}

func (h *Handler) RequestLogger(r *http.Request) *log.Logger {
	return log.FromContext(r.Context())
}

func (h *Handler) InternalServerError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	h.template.ExecuteTemplate(w, "not-found-error", nil)
}

func (h *Handler) CurrentUser(r *http.Request) (*users.User, error) {
	s := h.Session(r)
	userId, ok := s.Values[UserIDSessionKey]
	if !ok {
		return nil, ErrNoUsesr
	}

	user, err := h.userSvc.FindById(r.Context(), userId.(uint64))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
