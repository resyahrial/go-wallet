package deposit

import (
	"context"
	"time"

	"github.com/thoas/go-funk"
)

type DepositUsecaseInterface interface {
	Balance(context.Context, interface{}, interface{}) *Balance
	Threshold(context.Context, interface{}, interface{}) []*DepositRequest
}

type DepositUsecase struct{}

func New() DepositUsecaseInterface {
	return &DepositUsecase{}
}

func (u *DepositUsecase) Balance(ctx context.Context, currBalance interface{}, input interface{}) (newBalance *Balance) {
	var ok bool
	var req *DepositRequest
	if req, ok = input.(*DepositRequest); !ok {
		return
	}

	if currBalance == nil {
		return &Balance{
			WalletId: req.WalletId,
			Amount:   req.Amount,
		}
	}

	newBalance = currBalance.(*Balance)
	newBalance.Amount += req.Amount
	return newBalance
}

func (u *DepositUsecase) Threshold(ctx context.Context, collector interface{}, input interface{}) (newCollector []*DepositRequest) {
	if collector == nil {
		return
	}

	var ok bool
	var dInput *DepositRequest
	if newCollector, ok = collector.([]*DepositRequest); !ok {
		return
	}

	if dInput, ok = input.(*DepositRequest); !ok {
		return
	}

	newCollector = append(newCollector, dInput)
	prev2Minutes := dInput.InsertedAt.AsTime().Add(-2 * time.Minute)

	return funk.Reduce(newCollector, func(acc []*DepositRequest, depositRequest *DepositRequest) []*DepositRequest {
		if depositRequest.InsertedAt.AsTime().Before(prev2Minutes) {
			return acc
		}
		return append(acc, depositRequest)
	}, []*DepositRequest{}).([]*DepositRequest)
}
