package encode_test

import (
	"math/big"
	"testing"

	. "github.com/andriidski/abiencode-go/encode"
	. "github.com/andriidski/abiencode-go/utils"
)

func TestUint256ToBe(t *testing.T) {
	b := Uint256ToBe(big.NewInt(1))
	expected := "0x0000000000000000000000000000000000000000000000000000000000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint128ToBe(t *testing.T) {
	b := Uint128ToBe(big.NewInt(1))
	expected := "0x00000000000000000000000000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint64ToBe(t *testing.T) {
	b := Uint64ToBe(uint64(1))
	expected := "0x0000000000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint56ToBe(t *testing.T) {
	b := Uint56ToBe(uint64(1))
	expected := "0x00000000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint48ToBe(t *testing.T) {
	b := Uint48ToBe(uint64(1))
	expected := "0x000000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint40ToBe(t *testing.T) {
	b := Uint40ToBe(uint64(1))
	expected := "0x0000000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint32ToBe(t *testing.T) {
	b := Uint32ToBe(uint32(1))
	expected := "0x00000001"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint16ToBe(t *testing.T) {
	b := Uint16ToBe(uint16(0x03))
	expected := "0x0003"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestUint8ToBe(t *testing.T) {
	b := Uint8ToBe(uint8(1))
	expected := "0x01"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}
