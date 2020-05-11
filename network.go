package fake

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func IPv4(prefix ...int) string {
	n := 1 << 8
	count := len(prefix)
	if count < 1 || count > 4 {
		return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n))
	}

	blocks := make([]string, 4)
	for i := 0; i < count; i++ {
		blocks[i] = strconv.Itoa(prefix[i] % n)
	}
	for i := count; i < 4; i++ {
		blocks[i] = strconv.Itoa(rand.Intn(n))
	}
	return strings.Join(blocks, ".")
}

func IPv6() string {
	n := 1 << 16
	return fmt.Sprintf("2001:0000:%x:%x:%x:%x:%x:%x", rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n))
}

func Mac() string {
	n := 1 << 8
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n))
}

func MacHalf() string {
	n := 1 << 8
	return fmt.Sprintf("%02x:%02x:%02x", rand.Intn(n), rand.Intn(n), rand.Intn(n))
}
