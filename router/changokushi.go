package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChangokushiService interface {
	GetWeapons(context.Context) error
}

func ChangokushiRoute(service ChangokushiService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/weapons", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetWeapons(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})
	}
}
