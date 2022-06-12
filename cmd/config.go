package main

import (
	"context"
	"fmt"
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
	"gopkg.in/yaml.v2"
)

var config Configuration

func initialize(env string) {
	setConfig(env)
	initWorkers()
	initProcessor()
}

func setConfig(env string) {
	confFilePath := fmt.Sprintf("config/%s.yml", env)
	f, err := os.Open(confFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&config); err != nil {
		panic(err)
	}

	log.SetOutput(os.Stdout)
}

func initWorkers() {
	var err error
	if app.DepositEmitter, err = message.NewEmitter(message.EmitterOpts{
		Brokers:     config.GetBroker(),
		Stream:      config.Stream.Deposits,
		MessageType: new(deposit.DepositWrapper),
	}); err != nil {
		panic(err)
	}

	if app.BalanceViewer, err = message.NewViewer(message.ViewerOpts{
		Brokers:         config.GetBroker(),
		Table:           collectors.BalanceTable,
		MessageListType: new(deposit.BalanceWrapper),
	}); err != nil {
		panic(err)
	}

	if app.ThresholdViewer, err = message.NewViewer(message.ViewerOpts{
		Brokers:         config.GetBroker(),
		Table:           collectors.ThresholdTable,
		MessageListType: new(collectors.CollectorWrapper[*deposit.DepositRequest]),
	}); err != nil {
		panic(err)
	}
}

func initProcessor() {
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	balanceProcessor := injector.InitBalanceProcessor(message.CollectorOpts{
		Brokers:         config.GetBroker(),
		Stream:          config.Stream.Deposits,
		MessageType:     new(deposit.DepositWrapper),
		MessageListType: new(deposit.BalanceWrapper),
	})
	grp.Go(balanceProcessor.Run(ctx))

	thresholdProcessor := injector.InitBalanceProcessor(message.CollectorOpts{
		Brokers:         config.GetBroker(),
		Stream:          config.Stream.Deposits,
		MessageType:     new(deposit.DepositWrapper),
		MessageListType: new(collectors.CollectorWrapper[*deposit.DepositRequest]),
	})
	grp.Go(thresholdProcessor.Run(ctx))

	go app.BalanceViewer.Run(ctx)
	go app.ThresholdViewer.Run(ctx)

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
