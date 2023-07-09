package abi_test

import (
	"testing"

	"github.com/andriidski/abiencode-go/abi"
)

func TestEncodeNotPacked(t *testing.T) {
	typesList := [][]string{
		{
			"int16",
			"bytes1",
			"uint16",
			"string",
		},
	}

	values := [][]interface{}{
		{
			int16(-1),
			[1]byte{0x42},
			uint16(0x03),
			"Hello, world!",
		},
	}

	expectedEncodings := []string{
		"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff420000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000d48656c6c6f2c20776f726c642100000000000000000000000000000000000000",
	}

	for i, types := range typesList {
		encoding := abi.Encode(types, values[i])
		if encoding != expectedEncodings[i] {
			t.Errorf("expected %s, got %s", expectedEncodings[i], encoding)
		}
	}
}
