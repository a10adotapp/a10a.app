package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type FinschiaService interface {
	GetMillionArthurs(context.Context) error
	GetMillionArthursSummary(context.Context) error
}

func FinschiaRoute(service FinschiaService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/million-arthurs", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetMillionArthurs(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})

		router.Get("/million-arthurs-summary", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetMillionArthursSummary(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})
	}
}
