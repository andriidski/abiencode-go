package utils

import "errors"

var ErrInt256Overflow = errors.New("int256 value > int256 max")
var ErrInt256Underflow = errors.New("int256 value < int256 min")
var ErrInt128Overflow = errors.New("int128 > int128 max")
var ErrInt128Underflow = errors.New("int128 < int128 min")
