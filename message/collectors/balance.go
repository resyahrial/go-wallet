package collectors

import (
	"context"

	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
)

var (
	group        goka.Group = "balance"
	BalanceTable goka.Table = goka.GroupTable(group)
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
		p, err := message.NewProcessor(c.opts, group, func(ctx goka.Context, msg interface{}) {
			var messageList []deposit.DepositRequest
			if v := ctx.Value(); v != nil {
				messageList = v.([]deposit.DepositRequest)
			}

			messageList = c.usecase.Run(ctx.Context(), messageList, msg)
			ctx.SetValue(messageList)
		})
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
