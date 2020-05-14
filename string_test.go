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

func TestLowerChars(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := LowerChars(rand.Intn(10) + 1)
		for _, r := range str {
			if r < 'a' || r > 'z' {
				t.Errorf("TestLowerChars(), invalid character '%s', %s", string(r), str)
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

func TestUpperChars(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		str := UpperChars(rand.Intn(10) + 1)
		for _, r := range str {
			if r < 'A' || r > 'Z' {
				t.Errorf("UpperChars(), invalid character '%s', %s", string(r), str)
			}
		}
	}
}

func TestPicStr(t *testing.T) {
	arr := []string{"a", "b", "c"}

	exists := func(n string) bool {
		for _, r := range arr {
			if r == n {
				return true
			}
		}
		return false
	}

	n := PickStr(arr)
	if !exists(n) {
		t.Errorf("%s does not exists in %v", n, arr)
	}
}

func TestLowerChar(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		c := LowerChar()
		if c < string('a') || c > string('z') {
			t.Errorf("LowerChar() = %s is not lowercase", c)
		}
	}
}

func TestUpperChar(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		c := UpperChar()
		if c < string('A') || c > string('Z') {
			t.Errorf("UpperChar() = %s is not uppercase", c)
		}
	}
}

func TestParagraph(t *testing.T) {

}
