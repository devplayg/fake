package fake

import (
	"strconv"
)

const (
	AlphabetLower  = "abcdefghijklmnopqrstuvwxyz"
	AlphabetgUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digit          = "0123456789"
	//AlphabetCount  = len(AlphabetLower)
	Alphabet     = AlphabetLower + AlphabetgUpper
	Characters   = AlphabetgUpper + AlphabetLower + Digit
	specialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func NumStr(min, max int) string {
	return strconv.Itoa(intRange(min, max))
}

func Alpha(n int) string {
	return randString(Alphabet, n)
}

func String(n int) string {
	return randString(Characters, n)
}

func LowerCase(n int) string {
	return randString(AlphabetLower, n)
}

func UpperCase(n int) string {
	return randString(AlphabetgUpper, n)
}
