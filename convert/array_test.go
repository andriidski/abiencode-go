package convert_test

import (
	"math/big"
	"reflect"
	"testing"

	. "github.com/andriidski/abiencode-go/convert"
	"github.com/ethereum/go-ethereum/common"
)

func TestToUint8Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = uint8(1)
	items[1] = uint8(2)
	items[2] = uint8(3)

	expected := []uint8{1, 2, 3}
	converted := ToUint8Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToUint64Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = uint64(1)
	items[1] = uint64(2)
	items[2] = uint64(3)

	expected := []uint64{1, 2, 3}
	converted := ToUint64Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToUint256Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = big.NewInt(1)
	items[1] = big.NewInt(2)
	items[2] = big.NewInt(3)

	expected := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	converted := ToUint256Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToInt8Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = int8(1)
	items[1] = int8(2)
	items[2] = int8(3)

	expected := []int8{1, 2, 3}
	converted := ToInt8Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToInt64Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = int64(1)
	items[1] = int64(2)
	items[2] = int64(3)

	expected := []int64{1, 2, 3}
	converted := ToInt64Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToInt256Array(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = big.NewInt(1)
	items[1] = big.NewInt(2)
	items[2] = big.NewInt(3)

	expected := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	converted := ToInt256Array(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToBoolArray(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = true
	items[1] = false
	items[2] = true

	expected := []bool{true, false, true}
	converted := ToBoolArray(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToBytesArray(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = []byte{1}
	items[1] = []byte{2}
	items[2] = []byte{3}

	expected := [][]byte{{1}, {2}, {3}}
	converted := ToBytesArray(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToAddressArray(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = common.HexToAddress("0xf3958f7F68875735c24424D554eB46ddF8e8eD33")
	items[1] = common.HexToAddress("0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF")
	items[2] = common.HexToAddress("0x0000000000000000000000000000000000000000")

	expected := []common.Address{
		common.HexToAddress("0xf3958f7F68875735c24424D554eB46ddF8e8eD33"),
		common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff"),
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
	}
	converted := ToAddressArray(items)

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}

func TestToInterfaceArray(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = uint8(1)
	items[1] = uint8(2)
	items[2] = uint8(3)

	expected := []uint8{1, 2, 3}
	// Convert twice to test ToInterfaceArray.
	converted := ToUint8Array(ToInterfaceArray(items))

	if !reflect.DeepEqual(converted, expected) {
		t.Errorf("expected %v, got %v", expected, converted)
	}
}
