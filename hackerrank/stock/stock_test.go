package stock

import "testing"

func TestStockmax(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			"case1",
			[]int{5, 3, 2},
			0,
		}, {
			"case2",
			[]int{1, 2, 100},
			197,
		}, {
			"case3",
			[]int{1, 3, 1, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stockmax(tt.prices); got != tt.want {
				t.Errorf("Stockmax() = %v, want %v", got, tt.want)
			}
		})
	}
}
