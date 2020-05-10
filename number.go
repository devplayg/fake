package fake

import (
	"math"
	"math/big"
)

func Int8() int8 {
	return int8(intRange(math.MinInt8, math.MaxInt8))
}

func Int16() int16 {
	return int16(intRange(math.MinInt16, math.MaxInt16))
}

func Int32() int32 {
	return int32(intRange(math.MinInt32, math.MaxInt32))
}

func Int64() int64 {
	return bigIntRange(big.NewInt(math.MinInt64), big.NewInt(math.MaxInt64)).Int64()
}

func Uint8() uint8 {
	return uint8(intRange(0, math.MaxUint8))
}

func Uint16() uint16 {
	return uint16(intRange(0, math.MaxUint16))
}

func Uint32() uint32 {
	return uint32(intRange(0, math.MaxUint32))
}

func Uint64() uint64 {
	return bigIntRange(big.NewInt(0), new(big.Int).SetUint64(math.MaxInt64)).Uint64()
}

func Float32() float32 {
	return float32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

func Float32Range(min, max float32) float32 {
	return float32Range(min, max)
}

func Float64() float64 {
	return float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

func Float64Range(min, max float64) float64 {
	return float64Range(min, max)
}

func Num(min, max int) int {
	return int(bigIntRange(big.NewInt(int64(min)), big.NewInt(int64(max))).Int64())
}

func Num64(min, max int64) int64 {
	return bigIntRange(big.NewInt(int64(min)), big.NewInt(int64(max))).Int64()
}
