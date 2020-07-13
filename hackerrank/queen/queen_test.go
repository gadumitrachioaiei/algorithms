package queen

import "testing"

func TestCohabit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			"case3",
			3,
			false,
		},
		{
			"case4",
			4,
			true,
		}, {
			"case5",
			5,
			true,
		},
		{
			"case6",
			6,
			true,
		},
		{
			"case7",
			7,
			true,
		},
		{
			"case8",
			8,
			true,
		},
		{
			"case9",
			9,
			true,
		},
		{
			"case10",
			10,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cohabit(tt.n); got != tt.want {
				t.Errorf("Cohabit() = %v, want %v", got, tt.want)
			}
		})
	}
}
