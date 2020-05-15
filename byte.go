package fake

import "math/rand"

// Byte returns n size of byte slice
func Byte(n int) []byte {
	if n < 1 {
		return []byte{}
	}
	b := make([]byte, n)
	rand.Read(b)
	return b
}
