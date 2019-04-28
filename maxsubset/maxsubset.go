package maxsubset


// MaxSubsetLength returns the maximum number of elements from s, such that sum of any 2, is not divisible by k
func MaxSubsetLength(s []int, k int) int {
	max := 0
	for i := 0; i<len(s); i++ {
		// assume element i will be in the subset and check the result with this assumption
		if answer := maxSubsetLengthStart(s[i:i+1:i+1], s[i+1:], k); answer > max {
			max = answer
		}
	}
	return max
}

// maxSubsetLengthStart returns the number of elements of a + maximum number of elements from s, such that sum of any 2, is not divisible by k
// a is the existing subset
// s is the list from which we can choose elements
func maxSubsetLengthStart(a[]int, s []int, k int) int {
	var candidates [][]int //  a list of other candidates, that contain a, and that we should try
	var indexes []int // indexes is a list with an index from s, for each of the candidates, from which we can start choosing other elements
	for i, el := range s {
		isElOK := true
		for j := range a {
			if (a[j] + el) % k == 0 {
				isElOK = false
				if j > 0 {
					// remember the first elements that this combines fine with
					c := a[:j:j]
					c = append(c, el)
 					candidates = append(candidates, c)
					indexes = append(indexes, i)
				}
				break
			}
		}
		if isElOK {
			a = append(a, el)
		}
	}
	max := len(a)
	for i, l := range candidates {
		if answer := maxSubsetLengthStart(l, s[indexes[i] + 1:], k); answer > max {
			max = answer
		}
	}
	return max
}

