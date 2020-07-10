package candies

import "testing"

func TestCandies(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want int
	}{
		{
			"case1",
			[]int{1, 2, 3},
			6,
		}, {
			"case2",
			[]int{1, 2, 2},
			4,
		}, {
			"case3",
			[]int{1, 2, 1},
			4,
		}, {
			"case4",
			[]int{2, 1, 3},
			5,
		}, {
			"case5",
			[]int{2, 2, 3},
			4,
		}, {
			"case6",
			[]int{2, 1, 1},
			4,
		}, {
			"case7",
			[]int{3, 2, 1, 1, 2},
			9,
		}, {
			"case8",
			[]int{3, 2, 1, 1, 1, 2},
			10,
		}, {
			"case9",
			[]int{3, 2, 1, 1, 1, 2, 3},
			13,
		}, {
			"case10",
			[]int{1, 2, 3, 3},
			7,
		}, {
			"case11",
			[]int{1, 2, 3, 3, 3},
			8,
		}, {
			"case12",
			[]int{1, 2, 3, 3, 4},
			9,
		}, {
			"case13",
			[]int{1, 2, 3, 3, 3, 4},
			10,
		}, {
			"case14",
			[]int{1, 2, 3, 3, 3, 4, 5},
			13,
		}, {
			"case15",
			[]int{1, 2, 3, 3, 3, 4, 5, 5},
			14,
		}, {
			"case16",
			[]int{1, 2, 3, 3, 3, 2, 1},
			13,
		}, {
			"case17",
			[]int{1, 2, 3, 3, 3, 2, 1, 1, 2, 3},
			19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Candies(tt.a); got != tt.want {
				t.Errorf("Candies() = %v, want %v", got, tt.want)
			}
		})
	}
}
