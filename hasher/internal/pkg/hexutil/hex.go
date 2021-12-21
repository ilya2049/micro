package hexutil

import "encoding/hex"

func MustDecodeString(hexString string) []byte {
	bytes, _ := hex.DecodeString(hexString)

	return bytes
}
