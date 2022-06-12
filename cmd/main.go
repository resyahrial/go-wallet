package main

import (
	"log"
	"net/http"

	"github.com/resyarhial/go-wallet/cmd/injector"
)

func main() {
	r := injector.InitRouter()

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
