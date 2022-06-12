package deposit

import (
	"context"
)

type DepositUsecaseInterface[T DepositRequest] interface {
	Run(context.Context, []*T, interface{}) []*T
}

type DepositUsecase[T DepositRequest] struct{}

func New[T DepositRequest]() DepositUsecaseInterface[T] {
	return &DepositUsecase[T]{}
}

func (u *DepositUsecase[T]) Run(ctx context.Context, currBalance []*T, input interface{}) (newBalance []*T) {
	return newBalance
}
