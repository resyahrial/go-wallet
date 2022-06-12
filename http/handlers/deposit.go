package handlers

import (
	"net/http"
	"time"

	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	httputils "github.com/resyarhial/go-wallet/pkg/http-utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DepositHandlerInterface interface {
	Deposit(http.ResponseWriter, *http.Request)
}

type DepositHandler struct {
	emitter *goka.Emitter
}

func NewDepositHandler(emitter *goka.Emitter) DepositHandlerInterface {
	return &DepositHandler{emitter: emitter}
}

func (h *DepositHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var d deposit.DepositRequest
	if err := httputils.ReadRequestBody(r, &d); err != nil {
		panic(err)
	}
	d.InsertedAt = timestamppb.New(time.Now().UTC())
	if err := h.emitter.EmitSync(d.WalletId, &d); err != nil {
		panic(err)
	}
	httputils.WriteResponse(w, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "deposit has been processed",
	})
}
