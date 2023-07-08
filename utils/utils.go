package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ToHex(b []byte) string {
	return "0x" + common.Bytes2Hex(b)
}

func ToHexPadded(b []byte, nBytes int) string {
	return "0x" + common.Bytes2Hex(common.RightPadBytes(b, nBytes))
}

func BigIntFromString(str string) *big.Int {
	b, err := new(big.Int).SetString(str, 10)
	if !err {
		panic(err)
	}
	return b
}
