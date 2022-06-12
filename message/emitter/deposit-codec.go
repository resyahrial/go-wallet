package emitter

import (
	"github.com/resyarhial/go-wallet/internal/deposit"
)

type DepositCodec struct{}

func (c *DepositCodec) Encode(value interface{}) ([]byte, error) {
	return encode(value)
}

func (c *DepositCodec) Decode(data []byte) (interface{}, error) {
	var d deposit.DepositRequest
	return &d, decode(data, &d)
}
