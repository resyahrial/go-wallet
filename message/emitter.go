package message

import "github.com/lovoo/goka"

func NewEmitter(brokers []string, stream string, messageType goka.Codec) (*goka.Emitter, error) {
	return goka.NewEmitter(brokers, goka.Stream(stream), messageType)
}
