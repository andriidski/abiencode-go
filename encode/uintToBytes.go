package encode

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/andriidski/abiencode-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func Uint8ToBe(v uint8) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Uint16ToBe(v uint16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Uint32ToBe(v uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Uint40ToBe(v uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[3:]
}

func Uint48ToBe(v uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[2:]
}

func Uint56ToBe(v uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[1:]
}

func Uint64ToBe(v uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Uint128ToBe(v *big.Int) []byte {
	return common.LeftPadBytes(v.Bytes(), 16)
}

func Uint256ToBe(v *big.Int) []byte {
	b := new(big.Int)
	b = b.Set(v)

	if b.Sign() < 0 || b.BitLen() > 256 {
		b.And(b, utils.TwoToPow256)
	}

	return common.LeftPadBytes(b.Bytes(), 32)
}
