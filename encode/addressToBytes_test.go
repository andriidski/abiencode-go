package encode_test

import (
	"testing"

	. "github.com/andriidski/abiencode-go/encode"
	. "github.com/andriidski/abiencode-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func TestAddressToBytes(t *testing.T) {
	b := AddressToBytes(common.HexToAddress("0xf3958f7F68875735c24424D554eB46ddF8e8eD33"))
	expected := "0xf3958f7f68875735c24424d554eb46ddf8e8ed33"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
	b = AddressToBytes(common.HexToAddress("0x0000000000000000000000000000000000000000"))
	expected = "0x0000000000000000000000000000000000000000"
	if ToHex(b) != expected {
		t.Errorf("expected %s, got %s", expected, ToHex(b))
	}
}
