package encode

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/andriidski/abiencode-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func Int8ToBe(v int8) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Int16ToBe(v int16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Int32ToBe(v int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Int40ToBe(v int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[3:]
}

func Int48ToBe(v int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[2:]
}

func Int56ToBe(v int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()[1:]
}

func Int64ToBe(v int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return buf.Bytes()
}

func Int128ToBe(v *big.Int) []byte {
	b := new(big.Int)
	b = b.Set(v)

	if b.Sign() < 0 || b.BitLen() > 255 {
		b.And(b, utils.TwoToPow128m1)
	}
	return common.LeftPadBytes(v.Bytes(), 16)
}

// TODO: review
func Int256ToBe(v *big.Int) []byte {
	b := new(big.Int)
	b = b.Set(v)

	if b.Sign() < 0 || b.BitLen() > 255 {
		b.And(b, utils.TwoToPow256m1)
	}

	return common.LeftPadBytes(b.Bytes(), 32)
}
