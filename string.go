package fake

import (
	"math/rand"
	"strconv"
)

const (
	alphabetLower  = "abcdefghijklmnopqrstuvwxyz"
	alphabetgUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digit          = "0123456789"
	//AlphabetCount  = len(AlphabetLower)
	Alphabet     = alphabetLower + alphabetgUpper
	characters   = alphabetLower + alphabetgUpper + Digit
	specialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func NumStr(min, max int) string {
	return strconv.Itoa(intRange(min, max))
}

func Alpha(n int) string {
	return randString(Alphabet, n)
}

func String(n int) string {
	return randString(characters, n)
}

func LowerCase(n int) string {
	return randString(alphabetLower, n)
}

func UpperCase(n int) string {
	return randString(alphabetgUpper, n)
}

func Hex() string {
	return randHex()
}

func PickStr(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
