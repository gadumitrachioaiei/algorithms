// Package commonchild: https://www.hackerrank.com/challenges/common-child/problem
//
// Problem: A string is a child of another string, if it can be obtained from the second one by deleting character
// but preserving the order. Find the longest common child of two given strings.
//
// Solution: A character from the first string is part of the longest common child or not.
// Hence, make this solution recursively, with a cache for overlapping cases,
// and some optimizations for finding index from where to start looking.
package commonchild

func CommonChild(s1 string, s2 string) int {
	cache = NewCache(len(s1) + 1)
	indexes = make([][]int, 100)
	for i, c := range s2 {
		indexes[c] = append(indexes[c], i)
	}
	return commonChild(s1, s2, 0, 0)
}

func commonChild(s1 string, s2 string, s1Index, s2Index int) int {
	n := len(s1)
	if s1Index > n-1 || s2Index > n-1 {
		return 0
	}
	var result int
	// we handle the case when the first char is not in the common child
	// we should first handle this case, as it provides better cache hit
	if r := getAndSetCache(s1, s2, s1Index+1, s2Index); r > result {
		result = r
	}
	// we handle the case when the first char is in the common child
	charIndexes := indexes[int32(s1[s1Index])]
	for _, index := range charIndexes {
		if index >= s2Index {
			// we need to do calculations only if we can achieve a bigger result
			if result < n-index && result < n-s1Index {
				if r := 1 + getAndSetCache(s1, s2, s1Index+1, index+1); r > result {
					result = r
				}
			}
			// we cannot find a bigger child starting from a later point in s2
			break
		}
	}
	return result
}

func getAndSetCache(s1 string, s2 string, i, j int) int {
	v, ok := cache.Get(i, j)
	if ok {
		return v
	}
	v = commonChild(s1, s2, i, j)
	cache.Set(i, j, v)
	return v
}

// indexes represents the indexes where each character is present in the second string.
// It is made as an array, because map is not fast enough to pass the tests on hackerrank.
// The assumption is that the characters are in the range [A-Z].
var indexes [][]int

// cache is a cache for longest common child of two given substrings of the original strings.
// It represents a substring as having a startIndex up to the length of the original string.
var cache *Cache

type Element struct {
	startIndex1, startIndex2 int
}

type Cache struct {
	data [][]int
	hits int
}

func NewCache(cap int) *Cache {
	c := Cache{}
	c.data = make([][]int, cap)
	for i := 0; i < cap; i++ {
		c.data[i] = make([]int, cap)
		for j := 0; j < cap; j++ {
			c.data[i][j] = -1
		}
	}
	return &c
}

func (c *Cache) Set(i, j int, v int) {
	c.data[i][j] = v
}

func (c *Cache) Get(i, j int) (int, bool) {
	v := c.data[i][j]
	if v == -1 {
		return 0, false
	}
	c.hits++
	return v, true
}
