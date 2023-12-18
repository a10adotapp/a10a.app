package main

import (
	"net/http"

	"github.com/a10adotapp/a10a.app/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	http.ListenAndServe(":3000", router.NewRouter())
}
