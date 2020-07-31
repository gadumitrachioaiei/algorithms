package validstring

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"case1",
			"abc",
			"YES",
		},
		{
			"case2",
			"abcc",
			"YES",
		},
		{
			"case3",
			"abccc",
			"NO",
		},
		{
			"case4",
			"aabbcd",
			"NO",
		},
		{
			"case5",
			"aabbccddeefghi",
			"NO",
		},
		{
			"case6",
			"abcdefghhgfedecba",
			"YES",
		},
		{
			"case7",
			"aabbc",
			"YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
