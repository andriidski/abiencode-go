package encode_test

import (
	"testing"

	. "github.com/andriidski/abiencode-go/encode"
	. "github.com/andriidski/abiencode-go/utils"
)

func TestStringToBytes(t *testing.T) {
	b := StringToBytes("test")
	expected := "0x74657374"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
	b = StringToBytes("")
	expected = "0x"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}
