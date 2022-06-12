package deposit

import (
	"context"
)

type DepositUsecaseInterface interface {
	Balance(context.Context, []DepositRequest, interface{}) []DepositRequest
	Threshold(context.Context, []DepositRequest, interface{}) []DepositRequest
}

type DepositUsecase struct{}

func New() DepositUsecaseInterface {
	return &DepositUsecase{}
}

func (u *DepositUsecase) Balance(ctx context.Context, currBalance []DepositRequest, input interface{}) (newBalance []DepositRequest) {
	return newBalance
}

func (u *DepositUsecase) Threshold(ctx context.Context, currBalance []DepositRequest, input interface{}) (newBalance []DepositRequest) {
	return newBalance
}
