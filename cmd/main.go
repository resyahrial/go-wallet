package main

import (
	"log"
	"net/http"

	"github.com/resyarhial/go-wallet/cmd/app"
	"github.com/resyarhial/go-wallet/cmd/injector"
)

func main() {
	initConfig()
	initProcessor()
	defer app.DepositEmitter.Finish()

	r := injector.InitRouter(app.DepositEmitter)

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
