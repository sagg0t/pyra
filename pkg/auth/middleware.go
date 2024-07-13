package auth

import (
	"encoding/gob"
	"net/http"

	"pyra/pkg/log"
	"pyra/pkg/session"
)

type sessionKey string

const UserIDSessionKey string = "userid"

func init() {
	gob.Register(sessionKey(""))
}

func Authenticated(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.FromContext(r.Context())
		s := session.FromCtx(r.Context())

		if _, ok := s.Values[UserIDSessionKey]; ok {
			next(w, r)
		} else {
			l.Trace("unauthenticated access to a protected endpoint")
			// s.AddFlash("Please sign in first")

			http.Redirect(w, r, "/signIn", http.StatusFound)
		}
	})
}

func NotAuthenticated(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.FromContext(r.Context())
		s := session.FromCtx(r.Context())

		if _, ok := s.Values[UserIDSessionKey]; !ok {
			next(w, r)
		} else {
			// s.AddFlash("You're already signed in")

			l.Trace("authenticated access to a public endpoint")
			http.Redirect(w, r, "/foodProducts", http.StatusFound)
		}
	})
}

func IsAuthenticated(r *http.Request) bool {
	log := log.FromContext(r.Context())
	s := session.FromCtx(r.Context())

	userId, ok := s.Values[UserIDSessionKey]

	log.Debug("IsAuthenticated", "value", ok)
	if ok {
		log.Debug("userId", "id", userId.(string))
	}

	return ok
}
