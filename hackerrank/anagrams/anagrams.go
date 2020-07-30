// Package anagrams: https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
// Problem:
// Given a string s, return the number of pairs of substrings, that are each other's anagram.
//
// Solution:
// For all possible substring pairs, check if they are anagrams, but also use a cache for being faster

package anagrams

func Anagrams(s string) int {
	n := len(s)
	c := Cache{}
	var r int
	for i := 1; i < n+1; i++ {
		c.Reset()
		for j := 0; j < n-i+1; j++ {
			charCountS1 := getAndSetCache(j, s[j:j+i], c)
			for k := j + 1; k < n-i+1; k++ {
				charCountS2 := getAndSetCache(k, s[k:k+i], c)
				if isAnagram(charCountS1, charCountS2) {
					r++
				}
			}
		}
	}
	return r
}

type Cache struct {
	m map[int]map[int32]int
}

func (c *Cache) Reset() {
	c.m = make(map[int]map[int32]int)
}

func (c *Cache) Set(startIndex int, m map[int32]int) {
	c.m[startIndex] = m
}

func (c *Cache) Get(startIndex int) map[int32]int {
	return c.m[startIndex]
}

func getAndSetCache(startIndex int, s string, c Cache) map[int32]int {
	m := c.Get(startIndex)
	if len(m) == 0 {
		m = count(s)
		c.Set(startIndex, m)
	}
	return m
}

func count(s string) map[int32]int {
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
	}
	return m
}

func isAnagram(m1, m2 map[int32]int) bool {
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}
