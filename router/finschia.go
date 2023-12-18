package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type FinschiaService interface {
	GetMillionAuthurs() error
}

func FinschiaRoute(service FinschiaService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/million-arthurs", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetMillionAuthurs(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("content-type", "application/json")
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Header().Set("content-type", "application/json")
			w.Write([]byte(`{}`))
		})
	}
}
