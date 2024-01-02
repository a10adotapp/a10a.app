package finschia_test

import (
	"context"
	"testing"

	"github.com/a10adotapp/a10a.app/lib/db"
	"github.com/a10adotapp/a10a.app/service/finschia"
	"github.com/joho/godotenv"
)

func TestGetMillionArthurs(t *testing.T) {
	godotenv.Load("../../.env")

	entClient := db.NewEntClient()

	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	service := finschia.NewFinschiaService(entClient)

	ctx := context.Background()

	t.Run("success test", func(t *testing.T) {
		if err := service.GetMillionArthurs(ctx); err != nil {
			t.Error(err)
		}
	})
}
