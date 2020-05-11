package fake

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func IPv4(n ...int) string {
	count := len(n)
	if count < 1 || count > 4 {
		return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	}

	blocks := make([]string, 4)
	for i := 0; i < count; i++ {
		blocks[i] = strconv.Itoa(n[i] % 256)
	}
	for i := count; i < 4; i++ {
		blocks[i] = strconv.Itoa(rand.Intn(256))
	}
	return strings.Join(blocks, ".")

}
