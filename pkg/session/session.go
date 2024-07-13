package session

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const SessionCookieName = "pyra-session"

type ctxKey string

const sessionCtxKey ctxKey = "session"

var (
	key   = os.Getenv("SESSION_SECRET")
	store = sessions.NewCookieStore([]byte(key))
)

func init() {
	// store.Options.SameSite = http.SameSiteStrictMode
	store.MaxAge(3600 * 24)
	store.Options.HttpOnly = true
	store.Options.Secure = false
}

type Session struct {
	*sessions.Session
	req *http.Request
}

func NewSession(r *http.Request) (*Session, error) {
	s := &Session{req: r}
	err := s.extract()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Session) extract() error {
	session, err := store.Get(s.req, SessionCookieName)
	if err != nil {
		return err
	}

	s.Session = session

	return nil
}

func CtxWithSession(s *Session, ctx context.Context) context.Context {
	return context.WithValue(ctx, sessionCtxKey, s)
}

func FromContext(ctx context.Context) *Session {
	return ctx.Value(sessionCtxKey).(*Session)
}
