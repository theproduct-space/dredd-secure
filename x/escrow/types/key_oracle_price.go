package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OraclePriceKeyPrefix is the prefix to retrieve all OraclePrice
	OraclePriceKeyPrefix = "OraclePrice/value/"
)

// OraclePriceKey returns the store key to retrieve a OraclePrice from the index fields
func OraclePriceKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
