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
)

// Injectors from injector.go:

func InitRouter(emitter *goka.Emitter) *mux.Router {
	depositHandlerInterface := handlers.NewDepositHandler(emitter)
	muxRouter := router.New(depositHandlerInterface)
	return muxRouter
}