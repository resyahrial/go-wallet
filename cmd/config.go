package main

import (
	"flag"

	"github.com/resyarhial/go-wallet/cmd/app"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
)

var (
	broker = flag.String("broker", "localhost:9092", "boostrap Kafka broker")
)

func initConfig() {
	var err error
	if app.DepositEmitter, err = message.NewEmitter([]string{*broker}, "deposits", new(deposit.DepositWrapper)); err != nil {
		panic(err)
	}

}
