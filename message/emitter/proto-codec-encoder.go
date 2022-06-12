package emitter

import (
	"errors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func encode(value interface{}) ([]byte, error) {
	if v, ok := value.(protoreflect.ProtoMessage); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("invalid type: message value must be protobuf")
}

func decode(data []byte, protoMessage protoreflect.ProtoMessage) error {
	return proto.Unmarshal(data, protoMessage)
}
