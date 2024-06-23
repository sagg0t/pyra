package server

import (
	"net/http"
	"time"

	"github.com/olehvolynets/pyra/pkg/log"
)

type responseStatusCatcher struct {
	w          http.ResponseWriter
	statusCode int
}

func (c *responseStatusCatcher) Header() http.Header {
	return c.w.Header()
}

func (c *responseStatusCatcher) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

func (c *responseStatusCatcher) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.w.WriteHeader(statusCode)
}

func (c *responseStatusCatcher) StatusCode() int {
	if c.statusCode == 0 {
		return http.StatusOK
	}

	return c.statusCode
}

func Logger(logger *log.Logger, f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startT := time.Now()
		// slog.Info("started", "method", r.Method, "path", r.URL.Path)

		ww := responseStatusCatcher{w: w}
		f.ServeHTTP(&ww, r)

		endT := time.Now()
		took := endT.Sub(startT)

		logger.Info("req",
			"status", ww.StatusCode(),
			"method", r.Method,
			"path", r.URL.Path,
			"took", took)
	})
}
