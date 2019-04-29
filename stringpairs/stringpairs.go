package stringpairs

// Solution returns a string contained within s, that contains no pairs of AA, BB, CC
// s must contain only characters A, B, C
func Solution(a string) string {
	const outside = 'D'
	var previous uint8 = outside
	input := []byte(a)
	for i := 0; i<len(input); i++ {
		if input[i] == previous {
			input = append(input[:i-1], input[i+1:]...)
			// we set previous to be the char before the removed pair
			if i-1 > 0 {
				previous = input[i-2]
			} else {
				previous = outside
			}
			i = i-2
		} else {
			// move on
			previous = input[i]
		}
	}
	return string(input)
}

// Solution2 returns a string contained within s, that contains no pairs of AA, BB, CC
// s must contain only characters A, B, C
func Solution2(s string) string {
	const outside = 'D'
	var previous uint8 = outside
	for i := 0; i<len(s); i++ {
		if s[i] == previous {
			s = s[:i-1] + s[i+1:]
			// we set previous to be the char before the removed pair
			if i-1 > 0 {
				previous = s[i-2]
			} else {
				previous = outside
			}
			i = i-2
		} else {
			// move on
			previous = s[i]
		}
	}
	return s
}

