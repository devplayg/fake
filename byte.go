package fake

import "math/rand"

func Byte(n int) []byte {
	if n < 1 {
		return []byte{}
	}
	b := make([]byte, n)
	rand.Read(b)
	return b
}
