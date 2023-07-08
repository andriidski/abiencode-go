package encode

import "github.com/ethereum/go-ethereum/common"

func AddressToBytes(v common.Address) []byte {
	return v.Bytes()
}
