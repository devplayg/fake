package fake

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func randString(str string, n int) string {
	size := len(str)
	s := make([]rune, n)
	for i := range s {
		s[i] = rune(str[rand.Intn(size)])
	}
	return string(s)
}

func lowerChars(n int) []rune {
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		s[i] = lowerChar()
	}
	return s
}

func upperChars(n int) []rune {
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		s[i] = upperChar()
	}
	return s
}

func lowerChar() rune {
	return rune(rand.Intn(26)) + 'a'
}

func upperChar() rune {
	return rune(rand.Intn(26)) + 'A'
}

func randHex() string {
	return fmt.Sprintf("%02x", rand.Intn(1<<8))
}

func dateWithJitter(t time.Time, jitter time.Duration) time.Time {
	if jitter == 0 {
		return t
	}
	if rand.Intn(2) == 1 {
		return t.Add(time.Duration(rand.Int63n(jitter.Milliseconds())) * time.Millisecond)
	}
	return t.Add(time.Duration(rand.Int63n(jitter.Milliseconds())*-1) * time.Millisecond)
}
