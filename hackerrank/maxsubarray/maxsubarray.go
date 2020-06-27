package maxsubarray

// FindSubArray returns the maximum sum of a slice of arr
func FindSubArray(arr []int) int {
	var max int
	var s int
	maxNegative := arr[0]
	for i := 0; i < len(arr); i++ {
		// this if is here only to eliminate as the whole solution to all elements being negative
		if arr[i] > maxNegative {
			maxNegative = arr[i]
		}
		if s == 0 {
			if arr[i] > 0 {
				s += arr[i]
			}
			continue
		}
		if arr[i] < 0 {
			// we also decide now if the current sum is worth storing
			if s > max {
				max = s
			}
			s += arr[i]
			if s < 0 {
				// start again
				s = 0
			}
		} else {
			s += arr[i]
		}
	}
	if s > max {
		max = s
	}
	if max == 0 {
		return maxNegative
	}
	return max
}

// FindSubSet returns the maximum sum of a subset of arr
func FindSubSet(arr []int) int {
	var s int
	maxNegative := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			s += arr[i]
		} else if arr[i] > maxNegative {
			maxNegative = arr[i]
		}
	}
	if s == 0 {
		return maxNegative
	}
	return s
}
