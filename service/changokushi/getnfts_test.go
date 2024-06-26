package changokushi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/a10adotapp/a10a.app/lib/db"
	"github.com/a10adotapp/a10a.app/service/changokushi"
	"github.com/joho/godotenv"
)

func TestGetNFTs(t *testing.T) {
	godotenv.Load("../../.env")

	entClient := db.NewEntClient()

	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	service := changokushi.NewChangokushiService(entClient)

	ctx := context.Background()

	err := service.GetWeapons(ctx)

	fmt.Printf("err: %v\n", err)
}
