package encode

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	uint8Type, _       = abi.NewType("uint8", "uint8", nil)
	uint64Type, _      = abi.NewType("uint64", "uint64", nil)
	boolType, _        = abi.NewType("bool", "bool", nil)
	bytesType, _       = abi.NewType("bytes", "bytes", nil)
	bytes32Type, _     = abi.NewType("bytes32", "bytes32", nil)
	stringArrayType, _ = abi.NewType("string[]", "string[]", nil)
	uint32ArrayType, _ = abi.NewType("uint32[]", "uint32[]", nil)
	uint64ArrayType, _ = abi.NewType("uint64[]", "uint64[]", nil)

	contractCallWithTokenCustodianOnly = abi.Arguments{{Type: uint64Type}, {Type: bytesType}, {Type: stringArrayType}, {Type: uint32ArrayType}, {Type: uint64ArrayType}, {Type: bytes32Type}}
	contractCallWithTokenUPC           = abi.Arguments{{Type: bytesType}}
)
