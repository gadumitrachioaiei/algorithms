// Package substringdiff solves this problem:
// given two strings and an in k, find the longest "common" substring that differs in at most k characters.
//
// Solution:
// We search starting from first character of each string, and consider substrings of length greater than k.
// When we compare two substrings, we split the substring into parts of size k and remainder:
// the substrings of size k are compared using a cache and
// the substrings of length remainder are compared directly.
// We also advanced the length as much as possible:
// advancing with the difference between k and the current diff or
// starting from length greater than the last result.
package substringdiff

// SubstringDiff returns the length of the longest "common" substring
func SubstringDiff(k int, s1, s2 string) int {
	cache := NewCache(k, s1, s2)
	cache.prefill()
	n := len(s1)
	result := k
	// all pairs of substrings of length k have at most k different characters
	// we are now interested in comparing of substrings of length k+1 up to n
	length := k + 1
	for j := 0; j < n-k; j++ {
		for l := 0; l < n-k; l++ {
			for i := length; i < n+1; {
				if j > n-i || l > n-i {
					break
				}
				diff := cache.compareCache(j, l, i)
				if diff > k {
					break
				}
				if result < i {
					result = i
					// we will compare only bigger substrings from now on
					length = result + 1
				}
				// we want to advance as much as possible,
				// with the difference between the target and the current diff
				if m := min(i+k-diff, n-j, n-l); i < m {
					i = m
				} else {
					i++
				}
			}
		}
	}
	return result
}

func min(i, j, k int) int {
	m := i
	if j < m {
		m = j
	}
	if k < m {
		m = k
	}
	return m
}

func compare(s1, s2 string) int {
	k := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			k++
		}
	}
	return k
}

type Cache struct {
	k      int
	s1, s2 string

	data [][]int // cached data for substrings of length k, filled at the start
}

func NewCache(k int, s1, s2 string) *Cache {
	return &Cache{
		k:  k,
		s1: s1,
		s2: s2,
	}
}

func (c *Cache) Get(start1, start2 int) int {
	return c.data[start1][start2]
}

func (c *Cache) compareCache(start1, start2, length int) int {
	end1 := start1 + length
	end2 := start2 + length
	k := c.k
	// no caching for k == 0
	if k == 0 {
		if c.s1[start1:end1] == c.s2[start2:end2] {
			return 0
		}
		return c.k + 1
	}
	var result int
	// we calculate the result by splitting length into parts of size k, and we handle the remainder afterwards
	for i := 0; i <= length-k; i = i + k {
		r := c.Get(start1+i, start2+i)
		result += r
	}
	q := length / k
	if length != q*k && result < k+1 {
		r := compare(c.s1[start1+q*k:end1], c.s2[start2+q*k:end2])
		result += r
	}
	return result
}

func (c *Cache) prefill() {
	n := len(c.s1)
	k := c.k
	if k == 0 {
		return
	}
	c.data = make([][]int, n-k+1)
	for j := 0; j < n-k+1; j++ {
		c.data[j] = make([]int, n-k+1)
		for l := 0; l < n-k+1; l++ {
			r := compare(c.s1[j:j+k], c.s2[l:l+k])
			c.data[j][l] = r
		}
	}
}
