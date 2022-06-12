package emitter

import "github.com/lovoo/goka"

func New(brokers []string, stream string, messageType goka.Codec) (*goka.Emitter, error) {
	return goka.NewEmitter(brokers, goka.Stream(stream), messageType)
}
