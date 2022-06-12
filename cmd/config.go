package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/resyarhial/go-wallet/cmd/app"
	"github.com/resyarhial/go-wallet/cmd/injector"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
	"github.com/resyarhial/go-wallet/message/collectors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	broker        = flag.String("broker", "localhost:9092", "boostrap Kafka broker")
	depositStream = flag.String("depositStream", "deposits", "deposit stream")
)

func initConfig() {
	var err error
	if app.DepositEmitter, err = message.NewEmitter([]string{*broker}, *depositStream, new(deposit.DepositWrapper)); err != nil {
		panic(err)
	}
	if app.BalanceViewer, err = message.NewViewer(message.ViewerOpts{}); err != nil {
		panic(err)
	}
}

func initProcessor() {
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	balanceProcessor := injector.InitBalanceProcessor(message.CollectorOpts{
		Brokers:         []string{*broker},
		Stream:          *depositStream,
		MessageType:     new(deposit.DepositWrapper),
		MessageListType: new(collectors.CollectorWrapper[*deposit.DepositRequest]),
	})
	grp.Go(balanceProcessor.Run(ctx))

	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}
	cancel()
	if err := grp.Wait(); err != nil {
		log.Println(err)
	}
	log.Println("done")
}
