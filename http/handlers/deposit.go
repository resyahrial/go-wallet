package handlers

import (
	"net/http"

	"github.com/resyarhial/go-wallet/internal/models"
	httputils "github.com/resyarhial/go-wallet/pkg/http-utils"
)

type DepositHandlerInterface interface {
	Deposit(http.ResponseWriter, *http.Request)
}

type DepositHandler struct{}

func NewDepositHandler() DepositHandlerInterface {
	return &DepositHandler{}
}

func (h *DepositHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var deposit models.Deposit
	if err := httputils.ReadRequestBody(r, &deposit); err != nil {
		panic(err)
	}

	if err := httputils.WriteResponse(w, deposit); err != nil {
		panic(err)
	}
}
