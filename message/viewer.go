package message

import "github.com/lovoo/goka"

type ViewerOpts struct {
	Brokers         []string
	Table           goka.Table
	MessageListType goka.Codec
}

func NewViewer(opts ViewerOpts) (*goka.View, error) {
	return goka.NewView(opts.Brokers, opts.Table, opts.MessageListType)
}
