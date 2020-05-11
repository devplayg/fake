package fake

import (
	"net"
	"testing"
)

func TestIPv4(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		ip := IPv4()
		if net.ParseIP(ip) == nil {
			t.Errorf("invalid IPv4 address: %s", ip)
		}
	}
}

func TestIPv6(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		ip := IPv6()
		if net.ParseIP(ip) == nil {
			t.Errorf("invalid IPv6 address: %s", ip)
		}
	}
}

func TestMac(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		macAddr := Mac()
		_, err := net.ParseMAC(macAddr)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestMacHalf(t *testing.T) {
	for i := 0; i < testTryCount; i++ {
		macAddr := "00:00:00:" + MacHalf()
		_, err := net.ParseMAC(macAddr)
		if err != nil {
			t.Error(err)
		}
	}
}
