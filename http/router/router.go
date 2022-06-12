package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/resyarhial/go-wallet/http/handlers"
)

func New(
	depositHandler handlers.DepositHandlerInterface,
	balanceHandler handlers.BalanceHandlerInterface,
) (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/deposit", depositHandler.Deposit).Methods(http.MethodPost)
	router.HandleFunc("/details/:walletId", balanceHandler.GetBalance).Methods(http.MethodGet)
	return
}
