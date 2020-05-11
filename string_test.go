package fake

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestAlpha(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := Alpha(rand.Intn(10) + 1)
		for _, r := range str {
			if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
				t.Errorf("TestAlpha(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}

func TestHex(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := Hex()
		for _, r := range str {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') {
				t.Errorf("Hex(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}

func TestLowerCase(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := LowerCase(rand.Intn(10) + 1)
		for _, r := range str {
			if r < 'a' || r > 'z' {
				t.Errorf("LowerCase(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}

func TestNumStr(t *testing.T) {
	min := 10
	max := 100
	for i := 0; i < testTryCount; i++ {
		str := NumStr(min, max)
		n, err := strconv.Atoi(str)
		if err != nil {
			t.Errorf("NumStr(), parse error: %s", str)
		}
		if n < min || n > max {
			t.Errorf("NumStr() = %s, range: %d ~ %d", str, min, max)
		}
	}
}

func TestString(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := Alpha(rand.Intn(10) + 1)
		for _, r := range str {
			if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
				t.Errorf("TestAlpha(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}

func TestUpperCase(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := UpperCase(rand.Intn(10) + 1)
		for _, r := range str {
			if r < 'A' || r > 'Z' {
				t.Errorf("UpperCase(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}
