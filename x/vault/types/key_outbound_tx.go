package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OutboundTxKeyPrefix is the prefix to retrieve all OutboundTx
	// fixme when migrate, the previous key is  "OutboundTx/value/"
	OutboundTxKeyPrefix         = "OutboundTx/value/outboundtx-"
	OutboundTxProposalKeyPrefix = "OutboundTx/value/proposal-"
)

// OutboundTxKey returns the store key to retrieve a OutboundTx from the index fields
func OutboundTxKey(
	requestID string,
) []byte {
	var key []byte

	requestIDBytes := []byte(requestID)
	key = append(key, requestIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
