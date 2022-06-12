package deposit

import (
	"context"
)

type DepositUsecaseInterface interface {
	Run(context.Context, []DepositRequest, interface{}) []DepositRequest
}

type DepositUsecase struct{}

func New() DepositUsecaseInterface {
	return &DepositUsecase{}
}

func (u *DepositUsecase) Run(ctx context.Context, currBalance []DepositRequest, input interface{}) (newBalance []DepositRequest) {
	return newBalance
}
