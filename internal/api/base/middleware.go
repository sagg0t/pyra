package base

import (
	"net/http"

	"pyra/pkg/log"
	"pyra/pkg/session"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.FromContext(r.Context())
		s := session.FromContext(r.Context())

		if _, ok := s.Values[UserIDSessionKey]; ok {
			next.ServeHTTP(w, r)
		} else {
			l.Trace("unauthenticated access to a protected endpoint")
			// s.AddFlash("Please sign in first")

			http.Redirect(w, r, "/signIn", http.StatusFound)
		}
	})
}

func IsAuthenticated(r *http.Request) bool {
	log := log.FromContext(r.Context())
	s := session.FromContext(r.Context())

	userId, ok := s.Values[UserIDSessionKey]

	if ok {
		log.Debug("USER_ID", "id", userId.(uint64))
	}

	return ok
}
