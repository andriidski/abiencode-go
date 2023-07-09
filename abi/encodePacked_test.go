package abi_test

import (
	"testing"

	"github.com/andriidski/abiencode-go/abi"
)

func TestEncodePacked(t *testing.T) {
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
		"0xffff42000348656c6c6f2c20776f726c6421",
	}

	for i, types := range typesList {
		encoding := abi.EncodePacked(types, values[i])
		if encoding != expectedEncodings[i] {
			t.Errorf("expected %s, got %s", expectedEncodings[i], encoding)
		}
	}
}

func TestEncodePackedBytesArray(t *testing.T) {
	typesList := [][]string{
		{
			"bytes1",
			"bytes2",
			"bytes3",
		},
		{
			"string",
			"bytes16",
		},
	}

	values := [][]interface{}{
		{
			[1]byte{0x01},
			[2]byte{0x01, 0x02},
			[3]byte{0x01, 0x02, 0x03},
		},
		{
			"Hello, world!",
			[16]byte{0x01, 0x02, 0x03, 0x04, 0x01, 0x02, 0x03, 0x04, 0x01, 0x02, 0x03, 0x04, 0x01, 0x02, 0x03, 0x04},
		},
	}

	expectedEncodings := []string{
		"0x010102010203",
		"0x48656c6c6f2c20776f726c642101020304010203040102030401020304",
	}

	for i, types := range typesList {
		encoding := abi.EncodePacked(types, values[i])
		if encoding != expectedEncodings[i] {
			t.Errorf("expected %s, got %s", expectedEncodings[i], encoding)
		}
	}
}
