package encode_test

import (
	"testing"

	. "github.com/andriidski/abiencode-go/encode"
	. "github.com/andriidski/abiencode-go/utils"
)

func TestBoolToBytes(t *testing.T) {
	b := BoolToBytes(true)
	expected := "0x01"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
	b = BoolToBytes(false)
	expected = "0x00"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}
