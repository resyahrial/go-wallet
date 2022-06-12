package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	httputils "github.com/resyarhial/go-wallet/pkg/http-utils"
	"github.com/thoas/go-funk"
)

type BalanceHandlerInterface interface {
	GetBalance(http.ResponseWriter, *http.Request)
}

type BalanceOpts struct {
	BalanceViewer   *goka.View
	ThresholdViewer *goka.View
}

type BalanceHandler struct {
	balanceViewer   *goka.View
	thresholdViewer *goka.View
}

func NewBalanceHandler(opts BalanceOpts) BalanceHandlerInterface {
	return &BalanceHandler{balanceViewer: opts.BalanceViewer, thresholdViewer: opts.ThresholdViewer}
}

func (h *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	walletId := mux.Vars(r)["wallet_id"]
	valBalance, err := h.balanceViewer.Get(walletId)
	if err != nil {
		httputils.WriteResponse(w, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "internal server error",
		})
	}

	balance, ok := valBalance.(*deposit.Balance)
	if !ok {
		httputils.WriteResponse(w, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "internal server error",
		})
	}

	valThreshold, err := h.thresholdViewer.Get(walletId)
	if err != nil {
		httputils.WriteResponse(w, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "internal server error",
		})
	}

	threshold, ok := valThreshold.([]*deposit.DepositRequest)
	if !ok {
		httputils.WriteResponse(w, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "internal server error",
		})
	}

	httputils.WriteResponse(w, map[string]interface{}{
		"code": http.StatusOK,
		"data": &deposit.BalanceResponse{
			WalletId: balance.WalletId,
			Amount:   balance.Amount,
			AboveThreshold: funk.Reduce(threshold, func(acc float32, depReq *deposit.DepositRequest) float32 {
				return acc + depReq.Amount
			}, 0).(float32) > 10000,
		},
	})
}
