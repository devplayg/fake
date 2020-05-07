package fake

import (
	crand "crypto/rand"
	"github.com/davecgh/go-spew/spew"
	"math"
	"math/big"
	"math/rand"
)

func Bool() bool {
	return rand.Intn(2) == 1
}

func Number(min, max int) int {
	return intRange(min, max)
}

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
	return int64Range(math.MinInt64, math.MaxInt64)
}

func Uint8() uint8 {
	return uint8(intRange(0, math.MaxUint8))
}

func Uint16() uint16 {
	return uint16(intRange(0, math.MaxUint16))
}

func Uint32() uint32 {
	return uint32(intRange(0, math.MaxInt32))
}

func Uint64() uint64 {
	return uint64(rand.Int63n(math.MaxInt64))
}

func intRange(from, to int) int {
	return int(int64Range(int64(from), int64(to)))
}

func int64Range(from, to int64) int64 {
	min := big.NewInt(int64(math.MinInt64))
	max := big.NewInt(int64(math.MaxInt64))
	c := big.NewInt(0)
	n := c.Sub(max, min).Add(c, big.NewInt(1))
	randNum, err := crand.Int(crand.Reader, n)
	if err != nil {
		spew.Dump(from)
		spew.Dump(to)
		panic(err)
	}

	return randNum.Add(randNum, min).Int64()
}

//func IntRange(min, max int) int {
//	return intRange(min, max)
//}
//

//func intRange(min, max int) int {
//	if min >= max {
//		return min
//	}
//	return rand.Intn(max-min+1) + min
//}
