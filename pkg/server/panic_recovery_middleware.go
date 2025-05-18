package server

import (
	"net/http"
	"runtime/debug"

	"pyra/pkg/log"
)

func PanicRecovery(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		l := log.FromContext(ctx)

		defer func() {
			if r := recover(); r != nil {
				l.TraceContext(ctx, "recovering from panic", "error", r)
				if l.Enabled(ctx, log.LevelDebug) {
					l.DebugContext(ctx, string(debug.Stack()))
				}

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		f.ServeHTTP(w, r)
	})
}
