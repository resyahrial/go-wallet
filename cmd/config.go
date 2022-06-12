package main

import (
	"flag"

	"github.com/resyarhial/go-wallet/cmd/app"
	"github.com/resyarhial/go-wallet/message/emitter"
)

var (
	broker = flag.String("broker", "localhost:9092", "boostrap Kafka broker")
)

func initConfig() {
	var err error
	if app.DepositEmitter, err = emitter.New([]string{*broker}, "deposits", new(emitter.DepositCodec)); err != nil {
		panic(err)
	}

}
