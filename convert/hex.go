package convert

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/andriidski/abiencode-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func HexToBuffer(data string) *bytes.Buffer {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(b)
}

func HexToBufferPadded(data string, paddedToBytes int) *bytes.Buffer {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}
	b = common.LeftPadBytes(b, paddedToBytes)
	return bytes.NewBuffer(b)
}

func HexToUint8(data string) uint8 {
	var val uint8
	binary.Read(HexToBuffer(data), binary.BigEndian, &val)
	return val
}

func HexToUint16(data string) uint16 {
	var val uint16
	binary.Read(HexToBuffer(data), binary.BigEndian, &val)
	return val
}

func HexToUint32(data string) uint32 {
	var val uint32
	binary.Read(HexToBufferPadded(data, 4), binary.BigEndian, &val)
	return val
}

func HexToUint64(data string) uint64 {
	var val uint64
	binary.Read(HexToBufferPadded(data, 8), binary.BigEndian, &val)
	return val
}

func HexToInt8(data string) int8 {
	var val int8
	binary.Read(HexToBuffer(data), binary.BigEndian, &val)
	return val
}

func HexToInt16(data string) int16 {
	var val int16
	binary.Read(HexToBuffer(data), binary.BigEndian, &val)
	return val
}

func HexToInt32(data string) int32 {
	var val int32
	binary.Read(HexToBufferPadded(data, 4), binary.BigEndian, &val)
	return val
}

func HexToInt64(data string) int64 {
	var val int64
	binary.Read(HexToBufferPadded(data, 8), binary.BigEndian, &val)
	return val
}

func HexToBigInt(data string) *big.Int {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}
	return new(big.Int).SetBytes(b)
}

func HexToBigIntSigned(data string) *big.Int {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}

	bi := new(big.Int).SetBytes(b)

	// Enforce int256 bound.
	if bi.BitLen() > 255 {
		bi.Sub(bi, utils.TwoToPow256)
	}

	return bi
}

func HexToBool(data string) bool {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}
	return b[0] == 0x01
}

func HexToAddress(data string) common.Address {
	return common.HexToAddress(data)
}

func HexToBytes(data string) []byte {
	b, err := hexutil.Decode(data)
	if err != nil {
		panic(err)
	}
	return b
}

func HexToString(data string) string {
	return string(HexToBytes(data))
}
