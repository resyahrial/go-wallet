// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/http/handlers"
	"github.com/resyarhial/go-wallet/http/router"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
	"github.com/resyarhial/go-wallet/message/collectors"
)

// Injectors from injector.go:

func InitRouter(emitter *goka.Emitter, balanceOpts handlers.BalanceOpts) *mux.Router {
	depositHandlerInterface := handlers.NewDepositHandler(emitter)
	balanceHandlerInterface := handlers.NewBalanceHandler(balanceOpts)
	muxRouter := router.New(depositHandlerInterface, balanceHandlerInterface)
	return muxRouter
}

func InitBalanceProcessor(opts message.CollectorOpts) collectors.BalanceCollectorInterface {
	depositUsecaseInterface := deposit.New()
	balanceCollectorInterface := collectors.NewBalanceCollector(depositUsecaseInterface, opts)
	return balanceCollectorInterface
}
