package change

import (
	"testing"
)

func TestChange(t *testing.T) {
	type args struct {
		units int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{4, []int{1, 2, 3}},
			4,
		}, {
			"2",
			args{10, []int{2, 5, 3, 6}},
			5,
		}, {
			"3",
			args{3, []int{8, 3, 1, 2}},
			3,
		},
		{
			"4",
			args{69, []int{25, 27, 40, 38, 17, 2, 28, 23, 9, 43, 18, 49, 15, 24, 19, 11, 1, 39, 32, 16, 35, 30, 48, 34, 20, 3, 6, 13, 44}},
			101768,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Change(tt.args.units, tt.args.coins); got != tt.want {
				t.Errorf("Change() = %v, want %v", got, tt.want)
			}
		})
	}
}
