package encode_test

import (
	"math/big"
	"testing"

	. "github.com/andriidski/abiencode-go/encode"
	. "github.com/andriidski/abiencode-go/utils"
)

func TestInt256ToBe(t *testing.T) {
	// Test max int256.
	b := Int256ToBe(BigIntFromString("57896044618658097711785492504343953926634992332820282019728792003956564819967"))
	expected := "0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	// Test min int256.
	b = Int256ToBe(BigIntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968"))
	expected = "0x8000000000000000000000000000000000000000000000000000000000000000"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	b = Int256ToBe(BigIntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819967"))
	expected = "0x8000000000000000000000000000000000000000000000000000000000000001"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

// TODO: int128 needs fixing
func TestInt128ToBe(t *testing.T) {
	// Test max int128.
	b := Int128ToBe(BigIntFromString("170141183460469231731687303715884105727"))
	expected := "0x7fffffffffffffffffffffffffffffff"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	// Test min int128.
	b = Int128ToBe(BigIntFromString("-170141183460469231731687303715884105728"))
	expected = "0x80000000000000000000000000000000"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	b = Int128ToBe(BigIntFromString("-170141183460469231731687303715884105727"))
	expected = "0x80000000000000000000000000000001"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	// Test int128.
	b = Int128ToBe(big.NewInt(-1))
	expected = "0xffffffffffffffffffffffffffffffff"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt64ToBe(t *testing.T) {
	// Test max int64.
	b := Int64ToBe(int64(9223372036854775807))
	expected := "0x7fffffffffffffff"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	// Test min int64.
	b = Int64ToBe(int64(-9223372036854775808))
	expected = "0x8000000000000000"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	b = Int64ToBe(int64(-9223372036854775807))
	expected = "0x8000000000000001"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}

	// Test int64.
	b = Int64ToBe(int64(-1))
	expected = "0xffffffffffffffff"

	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt56ToBe(t *testing.T) {
	b := Int56ToBe(int64(-1))
	expected := "0xffffffffffffff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt48ToBe(t *testing.T) {
	b := Int48ToBe(int64(-1))
	expected := "0xffffffffffff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt40ToBe(t *testing.T) {
	b := Int40ToBe(int64(-1))
	expected := "0xffffffffff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt32ToBe(t *testing.T) {
	b := Int32ToBe(int32(-1))
	expected := "0xffffffff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt16ToBe(t *testing.T) {
	b := Int16ToBe(int16(-1))
	expected := "0xffff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}

func TestInt8ToBe(t *testing.T) {
	b := Int8ToBe(int8(-1))
	expected := "0xff"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}
