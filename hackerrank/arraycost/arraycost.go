// Package arraycost solves the folowwing problem:
// given an array B, make up another array A, such that 1<= A[i] <= B[i] and
// the sum of distances between neighbour elements ( named cost ) is maximised.
//
// Solution:
// the first element will either be min or maximum
// and recursively, for the rest of the elements, we calculate two costs:
// cost when the first element of the remaining ones is min, and when it is max.
// Then, we conclude the final cost:
// 1. when the first element is min, we associate it with the second cost
// 2. when the first element is max, we associate it with either the first or second cost
package arraycost

// Cost returns the cost of an array A
func Cost(B []int) int {
	c1, c2 := cost(B)
	if c1 < c2 {
		return c2
	}
	return c1
}

// cost returns two costs for an array A, as defined above
// 1. when A[0] = 1 ( first element is min )
// 2. when A[0] = B[0] ( first element is max )
func cost(B []int) (int, int) {
	if len(B) == 2 {
		return B[1] - 1, max(B[0]-1, abs(B[0]-B[1]))
	}
	minCost, maxCost := cost(B[1:])
	// when current el is B[0], we associate it with either previous min cost or max cost
	newMaxCost := max((B[0] - 1 + minCost), abs(B[0]-B[1])+maxCost)
	// when current element is 1, we associate it with previous max
	newMinCost := (B[1] - 1) + maxCost
	return newMinCost, newMaxCost
}

func max(a1, a2 int) int {
	if a1 < a2 {
		return a2
	}
	return a1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
