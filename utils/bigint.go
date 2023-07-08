package utils

import (
	"math/big"
)

var (
	TwoToPow256   = new(big.Int).Lsh(big.NewInt(1), 256)         // 2 ** 256
	TwoToPow256m1 = new(big.Int).Sub(TwoToPow256, big.NewInt(1)) // 2 ** 256 - 1

	TwoToPow128   = new(big.Int).Lsh(big.NewInt(1), 128)         // 2 ** 128
	TwoToPow128m1 = new(big.Int).Sub(TwoToPow128, big.NewInt(1)) // 2 ** 128 - 1

	TwoToPow127   = new(big.Int).Lsh(big.NewInt(1), 127)         // 2 ** 127
	TwoToPow127m1 = new(big.Int).Sub(TwoToPow127, big.NewInt(1)) // 2 ** 127 - 1

	TwoToPow255   = new(big.Int).Lsh(big.NewInt(1), 255)         // 2 ** 255
	TwoToPow255m1 = new(big.Int).Sub(TwoToPow255, big.NewInt(1)) // 2 ** 255 - 1
)
