//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/http/handlers"
	"github.com/resyarhial/go-wallet/http/router"
)

func InitRouter(emitter *goka.Emitter) *mux.Router {
	wire.Build(
		handlers.NewDepositHandler,
		router.New,
	)
	return nil
}
