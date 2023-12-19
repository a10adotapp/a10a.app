package middleware

import (
	"net/http"

	"github.com/a10adotapp/a10a.app/lib/log"
)

func LoggerInjectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.NewLogger()

		r = r.WithContext(log.NewContextWithLogger(r.Context(), logger))

		next.ServeHTTP(w, r)
	})
}
