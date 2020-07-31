// Package validstring: https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
package validstring

func isValid(s string) string {
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
	}
	// handle a special case, so we can reason easier later
	if oneZeroOthersEqual(s, m) {
		return "YES"
	}
	// we conclude:
	// for a string to be valid, two different characters:
	// 1. have equal values, in which case they represent the minimal value.
	// 2. have different values with a difference of 1
	firstChar, secondChar := s[0], s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			secondChar = s[i]
			break
		}
	}
	if secondChar == firstChar {
		return "YES"
	}
	countChar1 := m[int32(firstChar)]
	countChar2 := m[int32(secondChar)]
	diff := countChar1 - countChar2
	if diff > 1 || diff < -1 {
		return "NO"
	}
	if diff == 0 {
		min := countChar1
		over := 0
		for _, v := range m {
			d := v - min
			if d == 0 {
				continue
			}
			if d == 1 {
				over++
				if over > 1 {
					return "NO"
				}
				continue
			}
			return "NO"
		}
		return "YES"
	}
	min := countChar1
	if diff == 1 {
		min = countChar2
	}
	delete(m, int32(firstChar))
	delete(m, int32(secondChar))
	for _, v := range m {
		if v != min {
			return "NO"
		}
	}
	return "YES"
}

// oneZeroOthersEqual checks a special case:
// if one cracter is found once and all the other characters are found in equal numbers, greater than 1
func oneZeroOthersEqual(s string, m map[int32]int) bool {
	var countNon1, count1 int
	var valueNon1 int
	countEl0 := m[int32(s[0])]
	if countEl0 == 1 {
		count1 = 1
	} else {
		countNon1 = 1
		valueNon1 = countEl0
	}
	delete(m, int32(s[0]))
	defer func() {
		m[int32(s[0])] = countEl0
	}()
	for _, v := range m {
		if v == 1 {
			count1++
		} else {
			countNon1++
			if valueNon1 == 0 {
				valueNon1 = v
			}
		}
		if count1 > 1 {
			return false
		}
		if countNon1 > 1 && v != 1 && v != valueNon1 {
			return false
		}
	}
	return true
}
