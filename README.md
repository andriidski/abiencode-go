# abiencode-go 

`abiEncode()` and [`abiEncodePacked()`](https://docs.soliditylang.org/en/develop/abi-spec.html#strict-encoding-mode) in Go.

## Usage

#### `Encode`

Solidity:
```solidity
abi.encode(int16(-1), bytes1(0x42), uint16(0x03), string("Hello, world!"))
```

Go:
```go
hex := abi.Encode(
    []string{"int16","bytes1","uint16","string"}, 
    []interface{}{int16(-1),[]byte{0x42},uint16(0x03),"Hello, world!"},
)
```

Output:
```
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff420000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000d48656c6c6f2c20776f726c642100000000000000000000000000000000000000
```

#### `EncodePacked`

Solidity:
```solidity
abi.encodePacked(int16(-1), bytes1(0x42), uint16(0x03), string("Hello, world!"))
```

Go:
```go
hex := abi.EncodePacked(
    []string{"int16","bytes1","uint16","string"}, 
    []interface{}{int16(-1),[]byte{0x42},uint16(0x03),"Hello, world!"},
)
```

Output:
```
0xffff42000348656c6c6f2c20776f726c6421
```