package fake

import "math/rand"

// Bool returns a random bool
func Bool() bool {
	return rand.Intn(2) == 1
}
