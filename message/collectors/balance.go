package collectors

import (
	"context"

	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
)

var (
	group goka.Group = "balance"
	Table goka.Table = goka.GroupTable(group)
)

type BalanceCollectorInterface interface {
	Run(context.Context) func() error
}

type BalanceCollector struct {
	usecase deposit.DepositUsecaseInterface[deposit.DepositRequest]
	opts    message.CollectorOpts
}

func NewBalanceCollector(
	usecase deposit.DepositUsecaseInterface[deposit.DepositRequest],
	opts message.CollectorOpts,
) BalanceCollectorInterface {
	return &BalanceCollector{usecase: usecase, opts: opts}
}

func (c *BalanceCollector) Run(ctx context.Context) func() error {
	return func() error {
		p, err := message.NewProcessor(c.opts, group, callback[*deposit.DepositRequest](c.usecase))
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
