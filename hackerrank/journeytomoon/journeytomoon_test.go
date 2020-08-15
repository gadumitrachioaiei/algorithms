package journeytomoon

import "testing"

func TestJourneyToMoon(t *testing.T) {
	type args struct {
		n         int32
		astronaut [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"case1",
			args{5, [][]int32{{0, 1}, {2, 3}, {0, 4}}},
			6,
		},
		{
			"case2",
			args{4, [][]int32{{0, 2}}},
			5,
		},
		{
			"case3",
			args{100000, [][]int32{{1, 2}, {3, 4}}},
			4999949998,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := journeyToMoon(tt.args.n, tt.args.astronaut); got != tt.want {
				t.Errorf("JourneyToMoon() = %v, want %v", got, tt.want)
			}
		})
	}
}
