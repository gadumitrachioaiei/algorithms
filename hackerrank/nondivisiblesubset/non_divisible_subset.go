package nondivisiblesubset

// Find returns the length of the longest subset of s,
// such that the sum of any two of its numbers is not divisible by d
func Find(s []int, d int) int {
	// we split the subset into two parts, with remainders is less than d/2 and greater than two, respectively
	// after that we add/replace numbers in the first part with numbers from the second part,
	// based on the number of representatives of the mutually exclusive remainder
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	var c1, c2 int
	var c *int
	for i := range s {
		r := s[i] % d
		var m map[int]int
		if r <= int(d/2) {
			m = m1
			c = &c1
		} else {
			m = m2
			c = &c2
		}
		if v := m[r]; v == 0 {
			m[r] = 1
			*c++
		} else if (2*r)%d != 0 {
			m[r] = v + 1
			*c++
		}
	}

	// we choose m1 to have the biggest numbers of elements
	if c1 < c2 {
		m1, m2 = m2, m1
		c1, c2 = c2, c1
	}
	// we see if we can should elements from m2 to m1
	for k2, v2 := range m2 {
		if v := v2 - m1[d-k2]; v > 0 {
			c1 += v
		}
	}
	return c1
}
