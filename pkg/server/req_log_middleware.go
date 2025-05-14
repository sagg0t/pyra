package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"pyra/pkg/log"
)

type responseStatusCatcher struct {
	http.ResponseWriter
	statusCode int
}

func (c *responseStatusCatcher) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
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

		traceId := uuid.New()
		l := logger.With("traceId", traceId.String())

		ctx := r.Context()
		ctxWithLog := log.CtxWithLogger(ctx, l)
		r = r.WithContext(ctxWithLog)

		ww := responseStatusCatcher{ResponseWriter: w}
		f.ServeHTTP(&ww, r)

		endT := time.Now()
		took := endT.Sub(startT)

		var level log.Level
		if ww.StatusCode() < 500 {
			level = log.LevelInfo
		} else {
			level = log.LevelError
		}

		l.Log(ctx, level, "",
			"event", log.RequestEvent,
			"status", ww.StatusCode(),
			"method", r.Method,
			"path", r.URL.Path,
			"took", took,
			"location", ww.Header().Get("Location"),
		)
	})
}
