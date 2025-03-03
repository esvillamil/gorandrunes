package randrunes

import (
"math/rand"
)


// RandRunes generates a random string of runes.
//
// length is the desired length of the string.
// strInput is the set of runes to choose from.
// Returns a random string of runes.
func RandRunes (length int, strInput string) string {
	var letters = []rune(strInput)
	strOutput := make([]rune, length)
	for i := range length {
		strOutput[i] = letters[rand.Intn(len(letters))]
	}
return string(strOutput)
}

// RandRunesR generates a random string of runes.
//
// length is the desired length of the string.
// strInput is the set of runes to choose from.
// Returns a slice of runes of the specified length.
func RandRunesR (length int, strInput []rune) []rune {
	var letters = []rune(strInput)
	strOutput := make([]rune, length)
	for i := range length {
		strOutput[i] = letters[rand.Intn(len(letters))]
	}
return strOutput
}
