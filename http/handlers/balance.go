package handlers

import (
	"net/http"

	"github.com/lovoo/goka"
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
}
