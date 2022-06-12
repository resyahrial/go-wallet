package collectors

import (
	"context"

	"github.com/lovoo/goka"
	"github.com/resyarhial/go-wallet/internal/deposit"
	"github.com/resyarhial/go-wallet/message"
)

var (
	ThresholdGroup goka.Group = "threshold"
	ThresholdTable goka.Table = goka.GroupTable(ThresholdGroup)
)

type ThresholdCollectorInterface interface {
	Run(context.Context) func() error
}

type ThresholdCollector struct {
	usecase deposit.DepositUsecaseInterface
	opts    message.CollectorOpts
}

func NewThresholdCollector(
	usecase deposit.DepositUsecaseInterface,
	opts message.CollectorOpts,
) ThresholdCollectorInterface {
	return &ThresholdCollector{usecase: usecase, opts: opts}
}

func (c *ThresholdCollector) Run(ctx context.Context) func() error {
	return func() error {
		p, err := message.NewProcessor(c.opts, ThresholdGroup, func(ctx goka.Context, msg interface{}) {
			ctx.SetValue(c.usecase.Threshold(ctx.Context(), ctx.Value(), msg))
		})
		if err != nil {
			return err
		}
		return p.Run(ctx)
	}
}
