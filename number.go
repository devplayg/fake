package fake

import (
	"math"
	"math/big"
	"math/rand"
)

// Int8 returns a random int8
func Int8() int8 {
	return int8(intRange(math.MinInt8, math.MaxInt8))
}

// Int16 returns a random int16
func Int16() int16 {
	return int16(intRange(math.MinInt16, math.MaxInt16))
}

// Int32 returns a random int32
func Int32() int32 {
	return int32(intRange(math.MinInt32, math.MaxInt32))
}

// Int64 returns a random int64
func Int64() int64 {
	return bigIntRange(big.NewInt(math.MinInt64), big.NewInt(math.MaxInt64)).Int64()
}

// Uint8 returns a random uint8
func Uint8() uint8 {
	return uint8(intRange(0, math.MaxUint8))
}

// Uint16 returns a random uint16
func Uint16() uint16 {
	return uint16(intRange(0, math.MaxUint16))
}

// Uint32 returns a random uint32
func Uint32() uint32 {
	return uint32(intRange(0, math.MaxUint32))
}

// Uint64 returns a random uint64
func Uint64() uint64 {
	return bigIntRange(big.NewInt(0), new(big.Int).SetUint64(math.MaxInt64)).Uint64()
}

// Float32 returns a random float32
func Float32() float32 {
	return float32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float64 returns a random float64
func Float64() float64 {
	return float64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Int8Range returns a int8 between min and max
func Int8Range(min, max int8) int8 {
	return int8(intRange(int(min), int(max)))
}

// Int16Range returns a int16 between min and max
func Int16Range(min, max int16) int16 {
	return int16(intRange(int(min), int(max)))
}

// Int32Range returns a int32 between min and max
func Int32Range(min, max int32) int32 {
	return int32(intRange(int(min), int(max)))
}

// Int64Range returns a int64 between min and max
func Int64Range(min, max int64) int64 {
	return bigIntRange(big.NewInt(min), big.NewInt(max)).Int64()
}

// Uint8Range returns a uint8 between min and max
func Uint8Range(min, max uint8) uint8 {
	return uint8(intRange(int(min), int(max)))
}

// Uint16Range returns a uint16 between min and max
func Uint16Range(min, max uint16) uint16 {
	return uint16(intRange(int(min), int(max)))
}

// Uint32Range returns a uint32 between min and max
func Uint32Range(min, max uint32) uint32 {
	return uint32(intRange(int(min), int(max)))
}

// Uint64Range returns a uint64 between min and max
func Uint64Range(min, max uint64) uint64 {
	t1 := new(big.Int).SetUint64(min)
	t2 := new(big.Int).SetUint64(max)
	return bigIntRange(t1, t2).Uint64()
}

// Float32Range returns a float32 between min and max
func Float32Range(min, max float32) float32 {
	return float32Range(min, max)
}

// Float64Range returns a float64 between min and max
func Float64Range(min, max float64) float64 {
	return float64Range(min, max)
}

// Num returns a int between min and max
func Num(min, max int) int {
	return intRange(min, max)
}

// Num returns a int64 between min and max
func Num64(min, max int64) int64 {
	return int64Range(min, max)
}

// ShuffleInts shuffles a slice of ints
func ShuffleInts(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

// ShuffleStr shuffles a slice of strings
func ShuffleStr(arr []string) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

// PickInt returns one of the slice of ints
func PickInt(arr []int) int {
	return arr[rand.Intn(len(arr))]
}

// BigIntRange returns *big.Int between min and max
func BigIntRange(min, max *big.Int) *big.Int {
	return bigIntRange(min, max)
}

// Digit returns a number string
func Digit() string {
	return string(randDigit())
}

// NumCode returns a string with '#' converted to a number
func NumCode(str string) string {
	return sharpToNum(str)
}
