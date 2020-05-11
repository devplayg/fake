package fake

import "testing"

func BenchmarkIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4()
	}
}
