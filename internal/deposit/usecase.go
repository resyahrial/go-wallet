package deposit

import (
	"context"
	"time"

	"github.com/thoas/go-funk"
)

type DepositUsecaseInterface interface {
	Balance(context.Context, interface{}, interface{}) []*DepositRequest
	Threshold(context.Context, interface{}, interface{}) []*DepositRequest
}

type DepositUsecase struct{}

func New() DepositUsecaseInterface {
	return &DepositUsecase{}
}

func (u *DepositUsecase) Balance(ctx context.Context, currBalance interface{}, input interface{}) (newBalance []*DepositRequest) {
	if currBalance == nil {
		return
	}

	newBalance = currBalance.([]*DepositRequest)
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
