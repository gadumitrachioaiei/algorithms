// +build gofuzz

package coveringsubstring

import (
	"unicode/utf8"
)

// Fuzz should return 0 if the corpus should be ignored by the fuzzing tool, 1 otherwise
func Fuzz(data []byte) int {
	if isValid := utf8.Valid(data); !isValid {
		return 0
	}
	Substring(string(data))
	return 1
}
