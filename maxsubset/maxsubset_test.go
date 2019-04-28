package maxsubset

import "testing"

type testCase struct {
	s        []int
	k        int
	expected int
}

var tcs = []testCase{
	{[]int{19, 10, 12, 10, 24, 25, 22}, 4, 3},
	{[]int{2, 7, 1, 4}, 3, 3},
}

func TestMaxSubsetLength(t *testing.T) {
	for _, tc := range tcs {
		if answer := MaxSubsetLength(tc.s, tc.k); answer != tc.expected {
			t.Errorf("got max subset length: %d, expected: %d", answer, tc.expected)
		}
	}
}
