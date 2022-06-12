package message

import "github.com/lovoo/goka"

type EmitterOpts struct {
	Brokers     []string
	Stream      string
	MessageType goka.Codec
}

func NewEmitter(opts EmitterOpts) (*goka.Emitter, error) {
	return goka.NewEmitter(opts.Brokers, goka.Stream(opts.Stream), opts.MessageType)
}
