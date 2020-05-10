package fake

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
)

func intRange(min, max int) int {
	if min >= max {
		return min
	}
	return rand.Intn(max-min+1) + min
}

func int64Range(min, max int64) int64 {
	if min >= max {
		return min
	}
	return rand.Int63n(max-min+1) + min
}

func bigIntRange(min, max *big.Int) *big.Int {
	if min.Cmp(max) > 0 {
		return min
	}

	c := new(big.Int)
	c = c.Sub(max, min).Add(c, big.NewInt(1))
	randNum, err := crand.Int(crand.Reader, c)
	if err != nil {
		panic(err)
	}

	return randNum.Add(randNum, min)
}

func float32Range(min, max float32) float32 {
	if min == max {
		return min
	}
	return min + rand.Float32()*(max-min)
}

func float64Range(min, max float64) float64 {
	if min == max {
		return min
	}
	return rand.Float64()*(max-min) + min
}

func alpha() string {
	return string(Alphabet[rand.Intn(len(Alphabet))])
}
