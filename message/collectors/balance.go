package collectors

import (
	"context"

	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
)

var (
	BalanceGroup goka.Group = "balance"
	BalanceTable goka.Table = goka.GroupTable(BalanceGroup)
)

type BalanceCollectorInterface interface {
	Run(context.Context) func() error
}

type BalanceCollector struct {
	usecase deposit.DepositUsecaseInterface
	opts    message.CollectorOpts
}

func NewBalanceCollector(
	usecase deposit.DepositUsecaseInterface,
	opts message.CollectorOpts,
) BalanceCollectorInterface {
	return &BalanceCollector{usecase: usecase, opts: opts}
}

func (c *BalanceCollector) Run(ctx context.Context) func() error {
	return func() error {
		p, err := message.NewProcessor(c.opts, BalanceGroup, func(ctx goka.Context, msg interface{}) {
			ctx.SetValue(c.usecase.Balance(ctx.Context(), ctx.Value(), msg))
		})
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
