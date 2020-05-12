package fake

import (
	"testing"
)

func TestByte(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		if len(Byte(i)) != i {
			t.Error("err")
		}
	}
}
