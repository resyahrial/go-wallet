//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/http/handlers"
	"github.com/resyarhial/go-wallet/http/router"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
	"github.com/resyarhial/go-wallet/message/collectors"
)

func InitRouter(emitter *goka.Emitter) *mux.Router {
	wire.Build(
		handlers.NewDepositHandler,
		router.New,
	)
	return nil
}

func InitBalanceProcessor(opts message.CollectorOpts) collectors.BalanceCollectorInterface {
	wire.Build(
		deposit.New,
		collectors.NewBalanceCollector,
	)
	return nil
}
