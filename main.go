package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(10 * time.Second))

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

	http.ListenAndServe(":3000", router)
}
