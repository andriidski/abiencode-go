package abi

import (
	"math/big"

	"github.com/andriidski/abiencode-go/encode"
	"github.com/andriidski/abiencode-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

// EncodePacked is equivalent to abi.encodePacked(). The encoding is returned as a hex string.
//
// Example:
//
// Solidity call
//
//	abi.encodePacked(uint16(-1), uint16(0x03), "Hello, world!")
//
// Encode call
//
//	EncodePacked([]string{"uint16", "uint16", "string"}, []interface{}{uint16(-1), uint16(0x03), "Hello, world!"})
func EncodePacked(types []string, values []interface{}) string {
	data := []byte{}

	for i, _type := range types {
		encoding := EncodePackedValue(_type, values[i])
		data = append(data, encoding...)
	}

	return utils.ToHex(data)
}

// EncodePackedValue encodes a single value according to the given type.
func EncodePackedValue(_type string, value interface{}) []byte {
	switch _type {
	case "uint8":
		v, _ := value.(uint8)
		return encode.Uint8ToBe(v)
	case "uint16":
		v, _ := value.(uint16)
		return encode.Uint16ToBe(v)
	case "uint32":
		v, _ := value.(uint32)
		return encode.Uint32ToBe(v)
	case "uint40":
		v, _ := value.(uint64)
		return encode.Uint40ToBe(v)
	case "uint48":
		v, _ := value.(uint64)
		return encode.Uint48ToBe(v)
	case "uint56":
		v, _ := value.(uint64)
		return encode.Uint56ToBe(v)
	case "uint64":
		v, _ := value.(uint64)
		return encode.Uint64ToBe(v)
	case "uint72",
		"uint80",
		"uint88",
		"uint96",
		"uint104",
		"uint112",
		"uint120",
		"uint128":
		v, _ := value.(*big.Int)
		return encode.Uint128ToBe(v)
	case "uint136",
		"uint144",
		"uint152",
		"uint160",
		"uint168",
		"uint176",
		"uint184",
		"uint192",
		"uint200",
		"uint208",
		"uint216",
		"uint224",
		"uint232",
		"uint240",
		"uint248",
		"uint256":
		v, _ := value.(*big.Int)
		return encode.Uint256ToBe(v)
	case "int8":
		v, _ := value.(int8)
		return encode.Int8ToBe(v)
	case "int16":
		v, _ := value.(int16)
		return encode.Int16ToBe(v)
	case "int32":
		v, _ := value.(int32)
		return encode.Int32ToBe(v)
	case "int40":
		v, _ := value.(int64)
		return encode.Int40ToBe(v)
	case "int48":
		v, _ := value.(int64)
		return encode.Int48ToBe(v)
	case "int56":
		v, _ := value.(int64)
		return encode.Int56ToBe(v)
	case "int64":
		v, _ := value.(int64)
		return encode.Int64ToBe(v)
	case "int72",
		"int80",
		"int88",
		"int96",
		"int104",
		"int112",
		"int120",
		"int128":
		v, _ := value.(*big.Int)
		return encode.Int128ToBe(v)
	case "int136",
		"int144",
		"int152",
		"int160",
		"int168",
		"int176",
		"int184",
		"int192",
		"int200",
		"int208",
		"int216",
		"int224",
		"int232",
		"int240",
		"int248",
		"int256":
		v, _ := value.(*big.Int)
		return encode.Int256ToBe(v)
	case "bytes1":
		v, _ := value.([1]byte)
		return v[:]
	case "bytes2":
		v, _ := value.([2]byte)
		return v[:]
	case "bytes3":
		v, _ := value.([3]byte)
		return v[:]
	case "bytes4":
		v, _ := value.([4]byte)
		return v[:]
	case "bytes5":
		v, _ := value.([5]byte)
		return v[:]
	case "bytes6":
		v, _ := value.([6]byte)
		return v[:]
	case "bytes7":
		v, _ := value.([7]byte)
		return v[:]
	case "bytes8":
		v, _ := value.([8]byte)
		return v[:]
	case "bytes9":
		v, _ := value.([9]byte)
		return v[:]
	case "bytes10":
		v, _ := value.([10]byte)
		return v[:]
	case "bytes11":
		v, _ := value.([11]byte)
		return v[:]
	case "bytes12":
		v, _ := value.([12]byte)
		return v[:]
	case "bytes13":
		v, _ := value.([13]byte)
		return v[:]
	case "bytes14":
		v, _ := value.([14]byte)
		return v[:]
	case "bytes15":
		v, _ := value.([15]byte)
		return v[:]
	case "bytes16":
		v, _ := value.([16]byte)
		return v[:]
	case "bytes17":
		v, _ := value.([17]byte)
		return v[:]
	case "bytes18":
		v, _ := value.([18]byte)
		return v[:]
	case "bytes19":
		v, _ := value.([19]byte)
		return v[:]
	case "bytes20":
		v, _ := value.([20]byte)
		return v[:]
	case "bytes21":
		v, _ := value.([21]byte)
		return v[:]
	case "bytes22":
		v, _ := value.([22]byte)
		return v[:]
	case "bytes23":
		v, _ := value.([23]byte)
		return v[:]
	case "bytes24":
		v, _ := value.([24]byte)
		return v[:]
	case "bytes25":
		v, _ := value.([25]byte)
		return v[:]
	case "bytes26":
		v, _ := value.([26]byte)
		return v[:]
	case "bytes27":
		v, _ := value.([27]byte)
		return v[:]
	case "bytes28":
		v, _ := value.([28]byte)
		return v[:]
	case "bytes29":
		v, _ := value.([29]byte)
		return v[:]
	case "bytes30":
		v, _ := value.([30]byte)
		return v[:]
	case "bytes31":
		v, _ := value.([31]byte)
		return v[:]
	case "bytes32":
		v, _ := value.([32]byte)
		return v[:]
	case "bool":
		v, _ := value.(bool)
		return encode.BoolToBytes(v)
	case "address":
		v, _ := value.(common.Address)
		return encode.AddressToBytes(v)
	case "string":
		v, _ := value.(string)
		return encode.StringToBytes(v)
	case "bytes":
		v, _ := value.([]byte)
		return v
	default:
		panic("unsupported type: " + _type)
	}
}
