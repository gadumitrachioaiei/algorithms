package interview

import "testing"

func TestKnapsack(t *testing.T) {
	type args struct {
		s []int
		v []int
		c int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{[]int{1, 2, 3}, []int{2, 3, 9}, 10},
			want: 14,
		},
		{
			name: "case2",
			args: args{[]int{1, 2, 3}, []int{9, 2, 3}, 10},
			want: 14,
		},
		{
			name: "case3",
			args: args{[]int{5, 5, 5}, []int{9, 2, 3}, 10},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack(tt.args.s, tt.args.v, tt.args.c); got != tt.want {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
