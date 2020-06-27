package maxsubarray

import (
	"testing"
)

func TestFindSubArray(t *testing.T) {
	type args struct {
		arr []int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{1, 2, 3, 4}},
			10,
		}, {
			"case2",
			args{[]int{2, -1, 2, 3, 4, -5}},
			10,
		},
		{
			"case3",
			args{[]int{-2, -3, -1, -4}},
			-1,
		},
		{
			"case4",
			args{[]int{-2, -3, -1, -4}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubArray(tt.args.arr); got != tt.want {
				t.Errorf("FindSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
