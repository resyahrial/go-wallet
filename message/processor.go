package message

import "github.com/lovoo/goka"

type CollectorOpts struct {
	Brokers         []string
	Stream          string
	MessageType     goka.Codec
	MessageListType goka.Codec
}

func NewProcessor(opts CollectorOpts, group goka.Group, cb goka.ProcessCallback) (*goka.Processor, error) {
	return goka.NewProcessor(
		opts.Brokers,
		goka.DefineGroup(
			group,
			goka.Input(
				goka.Stream(opts.Stream),
				opts.MessageType,
				cb,
			),
			goka.Persist(opts.MessageListType),
		),
	)
}
