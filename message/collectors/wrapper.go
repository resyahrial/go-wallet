package collectors

import (
	"encoding/json"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type CollectorWrapper[T protoreflect.ProtoMessage] struct{}

func (c *CollectorWrapper[T]) Encode(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (c *CollectorWrapper[T]) Decode(data []byte) (interface{}, error) {
	var m []T
	err := json.Unmarshal(data, &m)
	return m, err
}
