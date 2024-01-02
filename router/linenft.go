package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type LINENFTService interface {
	GetActivities(context.Context) error
}

func LINENFTRoute(service LINENFTService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/activities", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetActivities(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})
	}
}
