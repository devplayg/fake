package fake

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// IPv4 returns a random IPv4 address
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

// IPv6 returns a random IPv6 address
func IPv6() string {
	n := 1 << 16
	return fmt.Sprintf("2001:0000:%x:%x:%x:%x:%x:%x", rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n))
}

// Mac returns a random mac address
func Mac() string {
	n := 1 << 8
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n), rand.Intn(n))
}

// MacHalf returns a random half-mac address
func MacHalf() string {
	n := 1 << 8
	return fmt.Sprintf("%02x:%02x:%02x", rand.Intn(n), rand.Intn(n), rand.Intn(n))
}
