package fake

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

var rander = rand.Reader

type _UUID [16]byte

// UUID returns a random (Version 4) UUID, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// https://github.com/google/uuid/blob/bd451584982ecf4ca5b1e5938cf168e17e30d837/version4.go#L34
func UUID() string {
	var uid _UUID
	_, err := io.ReadFull(rander, uid[:])
	if err != nil {
		return NumCode("########-####-####-####-############")
	}
	uid[6] = (uid[6] & 0x0f) | 0x40 // Version 4
	uid[8] = (uid[8] & 0x3f) | 0x80 // Variant is 10

	var buf [36]byte
	encodeHex(buf[:], uid)
	return string(buf[:])
}

// https://github.com/google/uuid/blob/bd451584982ecf4ca5b1e5938cf168e17e30d837/uuid.go#L179
func encodeHex(dst []byte, uuid _UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
