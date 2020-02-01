// Package coveringsubstring finds the smallest covering substring
package coveringsubstring

import (
	"fmt"
	"unicode/utf8"
)

// Substring returns the smallest covering substring of the given string
func Substring(s string) string {
	l := len(s)
	c := countUnique(s, -1) // how many unique characters are in the string
	ssLength := len([]rune(s)) // the length of smallest covering substring
	result := s // the covering substring

	// helper variables that we need in the sliding windows loops
	m := make(map[rune]int, c) // this represents all characters that were found at the end of a sliding window ( when innner loop finishes )
	finishIndex := 0 // index of last character that we added to the map, at the end of a sliding window
	var previousRune rune // start of the previous sliding window
	countRunes := 0 // count of runes that have been consumed at one point in time ( i.e.: len( m )
	var (
		r rune
		size int
	) // character, and size of a decoded character
	for i := 0; i<l; i += size {
		r, size = utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size < 2 {
			panic(fmt.Sprintf("not valid unicode char at pos: %d %d, %+q", i, l, s[i:]))
		}
		// if the number of characters after the current one is too small
		if len([]rune(s[i:])) < c {
			break
		}
		// on the current transversal, we no longer have the previous character
		if i > 0 {
			v, ok := m[previousRune]
			if ok {
				countRunes--
				if v == 1 {
					delete(m, previousRune)
				} else {
					m[previousRune] -= 1
				}
			} else { // it is not actually possible to get here
				panic(fmt.Sprintf("cannot find expected value in the map: %+q", previousRune))
			}
		}
		previousRune = r
		// loop through the sliding window starting with last character from previous window
		var (
			r1 rune
			size1 int
		)
		for j := finishIndex; j < l; j+=size1 {
			r1, size1 = utf8.DecodeRuneInString(s[j:])
			if r1 == utf8.RuneError && size1 < 2 {
				panic(fmt.Sprintf("not valid unicode char at pos: %d %d, %+q", j, l, s[j:]))
			}
			// if j is the finish index from a previous loop, we don't add the current char to the map, as it already was in the previous loop
			if i == 0 || j != finishIndex {
				m[r1] += 1
				countRunes++
			}
			// we need a smaller substring then the one we already have
			if countRunes >= ssLength {
				finishIndex = j
				break
			}
			if len(m) == c {
				finishIndex = j
				ssLength = countRunes
				result = s[i : j+size1]
				break
			}
			if j + size1 > l-1 {
				finishIndex = j
			}
		}
	}
	return result
}
// SubstringR returns the smallest covering substring of the given string
// it uses recursion
// assumes s is ascii
func SubstringR(s string) string {
	l := len(s)
	if l == 1 || l == 0 {
		return s
	}
	m := unique(s[:l-1])
	ssLength := len(m)
	var result string
	if _, ok := m[rune(s[l-1])]; !ok {
		ssLength++
		return startFromLast(s, ssLength)
	}
	result = SubstringR(s[:l-1])
	lResult := len(result)
	// it might be that a shorter covering substring contains the last char
	if temp := startFromLast(s, ssLength); len(temp) < lResult {
		result = temp
	}
	return result
}

func countUnique(s string, max int) int {
	if len(s) == 1 {
		return 1
	}
	m := make(map[rune]struct{})
	for _, r := range s {
		m[r] = struct{}{}
	}
	return len(m)
}

func unique(s string) map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, r := range s {
		m[r] = struct{}{}
	}
	return m
}

func startFromLast(s string, countUnique int) string {
	m := make(map[rune]struct{}, countUnique)
	for i := len(s) - 1; i > -1; i-- {
		m[rune(s[i])] = struct{}{}
		if len(m) == countUnique {
			return s[i:]
		}
	}
	panic("cannot find substring")
}
