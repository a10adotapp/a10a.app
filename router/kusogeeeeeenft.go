package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type KusogeeeeeeNFTService interface {
	GetNFTs(context.Context) error
}

func KusogeeeeeeNFTRoute(service KusogeeeeeeNFTService) func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/nfts", func(w http.ResponseWriter, r *http.Request) {
			if err := service.GetNFTs(r.Context()); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message":"%s"}`, err.Error())))
				return
			}

			w.Write([]byte(`{}`))
		})
	}
}
