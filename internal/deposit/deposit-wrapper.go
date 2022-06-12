package deposit

import (
	protoencoder "github.com/resyarhial/go-wallet/pkg/proto-encoder"
)

type DepositWrapper struct{}

func (c *DepositWrapper) Encode(value interface{}) ([]byte, error) {
	return protoencoder.Encode(value)
}

func (c *DepositWrapper) Decode(data []byte) (interface{}, error) {
	var d DepositRequest
	return &d, protoencoder.Decode(data, &d)
}
