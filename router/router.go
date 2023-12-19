package router

import (
	"net/http"
	"time"

	"github.com/a10adotapp/a10a.app/router/middleware"
	"github.com/a10adotapp/a10a.app/service/finschia"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(chimiddleware.Timeout(10 * time.Second))
	router.Use(middleware.LoggerInjectMiddleware)

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("{\"message\":\"method not allowed\"}"))
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\":\"not found\"}"))
	})

	router.Get("/b", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://bookmarks.a10a.app", http.StatusMovedPermanently)
	})

	router.Route("/finschia", FinschiaRoute(finschia.NewFinschiaService()))

	return router
}
