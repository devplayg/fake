package fake

import "math/rand"

func Bool() bool {
	return rand.Intn(2) == 1
}
