package convert

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ToInterfaceArray(val interface{}) []interface{} {
	return val.([]interface{})
}

func ToUint8Array(val interface{}) []uint8 {
	array := ToInterfaceArray(val)
	uint8Array := make([]uint8, len(array))
	for i, item := range array {
		uint8Array[i] = item.(uint8)
	}
	return uint8Array
}

func ToUint16Array(val interface{}) []uint16 {
	array := ToInterfaceArray(val)
	uint16Array := make([]uint16, len(array))
	for i, item := range array {
		uint16Array[i] = item.(uint16)
	}
	return uint16Array
}

func ToUint32Array(val interface{}) []uint32 {
	array := ToInterfaceArray(val)
	uint32Array := make([]uint32, len(array))
	for i, item := range array {
		uint32Array[i] = item.(uint32)
	}
	return uint32Array
}

func ToUint64Array(val interface{}) []uint64 {
	array := ToInterfaceArray(val)
	uint64Array := make([]uint64, len(array))
	for i, item := range array {
		uint64Array[i] = item.(uint64)
	}
	return uint64Array
}

func ToUint256Array(val interface{}) []*big.Int {
	array := ToInterfaceArray(val)
	uint256Array := make([]*big.Int, len(array))
	for i, item := range array {
		uint256Array[i] = item.(*big.Int)
	}
	return uint256Array
}

func ToInt8Array(val interface{}) []int8 {
	array := ToInterfaceArray(val)
	int8Array := make([]int8, len(array))
	for i, item := range array {
		int8Array[i] = item.(int8)
	}
	return int8Array
}

func ToInt16Array(val interface{}) []int16 {
	array := ToInterfaceArray(val)
	int16Array := make([]int16, len(array))
	for i, item := range array {
		int16Array[i] = item.(int16)
	}
	return int16Array
}

func ToInt32Array(val interface{}) []int32 {
	array := ToInterfaceArray(val)
	int32Array := make([]int32, len(array))
	for i, item := range array {
		int32Array[i] = item.(int32)
	}
	return int32Array
}

func ToInt64Array(val interface{}) []int64 {
	array := ToInterfaceArray(val)
	int64Array := make([]int64, len(array))
	for i, item := range array {
		int64Array[i] = item.(int64)
	}
	return int64Array
}

func ToInt256Array(val interface{}) []*big.Int {
	array := ToInterfaceArray(val)
	int256Array := make([]*big.Int, len(array))
	for i, item := range array {
		int256Array[i] = item.(*big.Int)
	}
	return int256Array
}

func ToBoolArray(val interface{}) []bool {
	array := ToInterfaceArray(val)
	boolArray := make([]bool, len(array))
	for i, item := range array {
		boolArray[i] = item.(bool)
	}
	return boolArray
}

func ToBytesArray(val interface{}) [][]byte {
	array := ToInterfaceArray(val)
	bytesArray := make([][]byte, len(array))
	for i, item := range array {
		bytesArray[i] = item.([]byte)
	}
	return bytesArray
}

func ToAddressArray(val interface{}) []common.Address {
	array := ToInterfaceArray(val)
	addressArray := make([]common.Address, len(array))
	for i, item := range array {
		addressArray[i] = item.(common.Address)
	}
	return addressArray
}
