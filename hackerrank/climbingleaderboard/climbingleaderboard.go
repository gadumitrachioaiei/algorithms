package climbingleaderboard

// climbingLeaderboard returns the leader board after every game played
//  we return a list of length len(alice)
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	res := make([]int32, len(alice))
	// make up a map of indexes to rank
	m := map[int]int{0: 1}
	for i := 1; i < len(scores); i++ {
		if scores[i] == scores[i-1] {
			m[i] = m[i-1]
		} else {
			m[i] = m[i-1] + 1
		}
	}
	for i := 0; i < len(alice); i++ {
		index := binarySearch(scores, alice[i])
		if index == len(scores) {
			res[i] = int32(m[index-1] + 1)
		} else {
			res[i] = int32(m[index])
		}
	}
	return res
}

// binarySearch returns the index where the element would be inserted in scores
// or, in case el would be smaller then all elements in scores, len(scores)
func binarySearch(scores []int32, el int32) int {
	if len(scores) == 0 {
		return 0
	}
	i := len(scores) / 2
	if el == scores[i] {
		return i
	} else if el < scores[i] {
		// start looking at index i +1 to n
		if len(scores) == 1 {
			return 1
		}
		return i + 1 + binarySearch(scores[i+1:], el)
	} else {
		if len(scores) == 1 {
			return 0
		}
		return binarySearch(scores[:i], el)
	}
}
