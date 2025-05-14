package server

import (
	"net/http"

	"pyra/pkg/log"
)

func PanicRecovery(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		l := log.FromContext(ctx)

		defer func() {
			if r := recover(); r != nil {
				l.TraceContext(ctx, "recovering from panic", "error", r)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		f.ServeHTTP(w, r)
	})
}
