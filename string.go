package fake

import "math/rand"

const (
	AlphaLower    = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetCount = len(AlphaLower)
	Alphabet      = AlphaLower + AlphaUpper
	NumberStr     = "0123456789"
	specialChars  = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	//charset      = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
)

func Char() string {
	return lowercaseN(3)
}

func lowercase() string {
	return string(rune(rand.Intn(AlphabetCount) + 'a'))
}

func lowercaseN(count int) string {
	s := make([]rune, count)
	for i := range s {
		s[i] = rune(rand.Intn(AlphabetCount) + 'A')
	}
	return string(s)
}

func uppercase() string {
	return string(rune(rand.Intn(AlphabetCount) + 'a'))
}

func uppercaseN(count int) string {
	s := make([]rune, count)
	for i := range s {
		s[i] = rune(rand.Intn(AlphabetCount) + 'A')
	}
	return string(s)
}
