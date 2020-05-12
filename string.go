package fake

import (
	"github.com/devplayg/fake/data"
	"math/rand"
	"strconv"
	"unicode"
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

func Word() string {
	return data.LoremIpsum[rand.Intn(len(data.LoremIpsum))]
}

func Phrase() string {
	var str string
	n := len(data.LoremIpsum)
	count := intRange(1, 4)
	for i := 0; i < count; i++ {
		str += data.LoremIpsum[rand.Intn(n)]
		if i < (count - 1) {
			str += " "
		}
	}

	return str
}
func Sentence() string {
	var sentense string
	count := intRange(4, 10)
	for i := 0; i < count; i++ {
		sentense += Phrase()
		if i < (count - 1) {
			if rand.Intn(2) == 1 {
				sentense += ", "
			}
		}
	}
	return string(unicode.ToUpper(rune(sentense[0]))) + sentense[1:] + "."
}
func Paragraph() string {
	var paragraph string
	count := intRange(5, 8)
	for i := 0; i < count; i++ {
		paragraph += Sentence()
		if i < (count - 1) {
			paragraph += " "
		}
	}
	return paragraph
}
