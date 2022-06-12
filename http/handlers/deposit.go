package handlers

import (
	"net/http"
	"time"

	"github.com/resyarhial/go-wallet/internal/deposit"
	httputils "github.com/resyarhial/go-wallet/pkg/http-utils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DepositHandlerInterface interface {
	Deposit(http.ResponseWriter, *http.Request)
}

type DepositHandler struct{}

func NewDepositHandler() DepositHandlerInterface {
	return &DepositHandler{}
}

func (h *DepositHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var d deposit.DepositRequest
	if err := httputils.ReadRequestBody(r, &d); err != nil {
		panic(err)
	}
	d.InsertedAt = timestamppb.New(time.Now().UTC())

	var unmarshal deposit.DepositRequest
	if res, err := proto.Marshal(&d); err != nil {
		panic(err)
	} else if err = proto.Unmarshal(res, &unmarshal); err != nil {
		panic(err)
	} else if err = httputils.WriteResponse(w, &unmarshal); err != nil {
		panic(err)
	}
}
