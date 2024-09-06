package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SanmeiHoikuenService interface {
	ExplorePosts(ctx context.Context) error
}

func SanmeiHoikuenRoute(service SanmeiHoikuenService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/explore-posts", func(w http.ResponseWriter, r *http.Request) {
			if err := service.ExplorePosts(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})
	}
}
