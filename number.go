package fake

import (
	"math"
	"math/big"
	"math/rand"
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

func Float64() float64 {
	return float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Range
func Int8Range(min, max int8) int8 {
	return int8(intRange(int(min), int(max)))
}

func Int16Range(min, max int16) int16 {
	return int16(intRange(int(min), int(max)))
}

func Int32Range(min, max int32) int32 {
	return int32(intRange(int(min), int(max)))
}

func Int64Range(min, max int64) int64 {
	return bigIntRange(big.NewInt(min), big.NewInt(max)).Int64()
}

func Uint8Range(min, max uint8) uint8 {
	return uint8(intRange(int(min), int(max)))
}

func Uint16Range(min, max uint16) uint16 {
	return uint16(intRange(int(min), int(max)))
}

func Uint32Range(min, max uint32) uint32 {
	return uint32(intRange(int(min), int(max)))
}

func Uint64Range(min, max uint64) uint64 {
	t1 := new(big.Int).SetUint64(min)
	t2 := new(big.Int).SetUint64(max)
	return bigIntRange(t1, t2).Uint64()
}

func Float32Range(min, max float32) float32 {
	return float32Range(min, max)
}

func Float64Range(min, max float64) float64 {
	return float64Range(min, max)
}

// Utils
func Num(min, max int) int {
	return intRange(min, max)
}

func Num64(min, max int64) int64 {
	return int64Range(min, max)
}

func ShuffleInts(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

func ShuffleStr(arr []string) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

func PickInt(arr []int) int {
	return arr[rand.Intn(len(arr))]
}
