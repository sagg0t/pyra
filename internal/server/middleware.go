package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
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

func Logger(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startT := time.Now()
		// slog.Info("started", "method", r.Method, "path", r.URL.Path)

		ww := responseStatusCatcher{w: w}
		f.ServeHTTP(&ww, r)

		endT := time.Now()
		took := endT.Sub(startT).Seconds()

		slog.Info("req",
			"status", ww.StatusCode(),
			"method", r.Method,
			"path", r.URL.Path,
			"took", fmt.Sprintf("%.6f", took))
	})
}
