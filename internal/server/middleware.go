package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Logger(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startT := time.Now()
		// slog.Info("started", "method", r.Method, "path", r.URL.Path)

		f.ServeHTTP(w, r)

		endT := time.Now()
		took := endT.Sub(startT).Seconds()

		slog.Info("req",
			"method", r.Method,
			"path", r.URL.Path,
			"took", fmt.Sprintf("%.6f", took))
	})
}
