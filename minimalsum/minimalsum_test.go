package minimalsum

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0, 1, 8},
			4,
		},
		{
			"case2",
			args{4, 4, 0},
			0,
		},
		{
			"case3",
			args{6, 2, 0},
			0,
		},
		{
			"case4",
			args{6, 1, 1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimalSum(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
