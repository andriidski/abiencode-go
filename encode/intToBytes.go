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
	b := new(big.Int).Set(v)
	b = b.Set(v)

	// Check range.
	if b.Cmp(utils.TwoToPow127m1) > 0 {
		panic(utils.ErrInt128Overflow)
	}
	if b.Cmp(new(big.Int).Neg(utils.TwoToPow127)) < 0 {
		panic(utils.ErrInt128Underflow)
	}

	if b.Sign() < 0 {
		b.And(b, utils.TwoToPow128m1)
	}

	return common.LeftPadBytes(b.Bytes(), 16)
}

func Int256ToBe(v *big.Int) []byte {
	b := new(big.Int)
	b = b.Set(v)

	// Check range.
	if b.Cmp(utils.TwoToPow255m1) > 0 {
		panic(utils.ErrInt256Overflow)
	}
	if b.Cmp(new(big.Int).Neg(utils.TwoToPow255)) < 0 {
		panic(utils.ErrInt256Underflow)
	}

	if b.Sign() < 0 {
		b.And(b, utils.TwoToPow256m1)
	}

	return common.LeftPadBytes(b.Bytes(), 32)
}
