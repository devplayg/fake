package fake

import (
	"math/rand"
	"testing"
)

func TestByte(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		n := rand.Intn(10)

		if len(Byte(n)) != n {
			t.Error("err")
		}
	}
}
