package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/a10adotapp/a10a.app/lib/db"
	"github.com/a10adotapp/a10a.app/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})))

	entClient := db.NewEntClient()

	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		slog.Error("main", slog.Any("error", err))
	}

	http.ListenAndServe(":3000", router.NewRouter(entClient))
}
