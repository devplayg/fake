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
	digits         = "0123456789"
	alphabet       = alphabetLower + alphabetgUpper
	characters     = alphabetLower + alphabetgUpper + digits
	specialChars   = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// NumStr returns a number between min and max as a string.
func NumStr(min, max int) string {
	return strconv.Itoa(intRange(min, max))
}

// Alpha returns random alphabetical characters
func Alpha(n int) string {
	return randString(alphabet, n)
}

// String returns random characters
func String(n int) string {
	return randString(characters, n)
}

// LowerChar returns a random lowercase alphabet
func LowerChar() string {
	return string(lowerChar())
}

// LowerChars returns random lowercase alphabets
func LowerChars(n int) string {
	if n < 1 {
		return ""
	}
	return string(lowerChars(n))
}

// UpperChar returns a random uppercase alphabet
func UpperChar() string {
	return string(upperChar())
}

// UpperChars returns random uppercase alphabets
func UpperChars(n int) string {
	if n < 1 {
		return ""
	}
	return string(upperChars(n))
}

// Hex returns a random hex string
func Hex() string {
	return randHex()
}

// PicStr returns one of string in slice
func PickStr(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

// Word returns a random word
func Word() string {
	return data.LoremIpsum[rand.Intn(len(data.LoremIpsum))]
}

// Word returns a random phrase
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

// Word returns a random sentence
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

// Word returns a random paragraph
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
