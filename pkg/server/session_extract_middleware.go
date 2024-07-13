package server

import (
	"net/http"

	"pyra/pkg/log"
	"pyra/pkg/session"
)

type responseWithSession struct {
	http.ResponseWriter
	req            *http.Request
	sessionWritten bool
}

func (r *responseWithSession) WriteHeader(code int) {
	if !r.sessionWritten {
		r.writeSession()
	}

	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWithSession) Write(b []byte) (int, error) {
	if !r.sessionWritten {
		r.writeSession()
	}

	return r.ResponseWriter.Write(b)
}

func (r *responseWithSession) writeSession() {
	l := log.FromContext(r.req.Context())
	s := session.FromCtx(r.req.Context())

	err := s.Save(r.req, r.ResponseWriter)
	if err != nil {
		l.Error("session write", "error", err)
		panic(err)
	}

	r.sessionWritten = true
}

func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.FromContext(r.Context())

		s, err := session.Get(r)
		if err != nil {
			l.Error("session read", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		ctxWithSession := session.CtxWithSession(s, r.Context())
		r = r.WithContext(ctxWithSession)

		ww := &responseWithSession{
			req:            r,
			ResponseWriter: w,
		}

		next.ServeHTTP(ww, r)
	})
}
