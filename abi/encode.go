package abi

import (
	"github.com/andriidski/abiencode-go/utils"
	abi_geth "github.com/ethereum/go-ethereum/accounts/abi"
)

// Encode is equivalent to abi.encode(). The encoding is returned as a hex string.
//
// Example:
//
// Solidity call
//
//	abi.encode(uint16(-1), uint16(0x03), "Hello, world!")
//
// Encode call
//
//	Encode([]string{"uint16", "uint16", "string"}, []interface{}{uint16(-1), uint16(0x03), "Hello, world!"})
func Encode(types []string, values []interface{}) string {
	args := make(abi_geth.Arguments, 0)

	for _, _type := range types {

		arg := abi_geth.Argument{}
		var err error
		arg.Type, err = abi_geth.NewType(_type, "", nil)
		if err != nil {
			panic(err)
		}
		args = append(args, arg)
	}

	encoding, err := args.PackValues(values)
	if err != nil {
		panic(err)
	}
	return utils.ToHex(encoding)
}
