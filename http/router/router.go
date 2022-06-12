package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/resyarhial/go-wallet/http/handlers"
)

func New(
	depositHandler handlers.DepositHandlerInterface,
) (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/deposit", depositHandler.Deposit).Methods(http.MethodPost)
	return
}
