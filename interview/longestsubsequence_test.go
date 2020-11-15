package interview

import "testing"

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"case1",
			"masina",
			"ain",
		},
		{
			"case2",
			"pesedist",
			"eeist",
		},
		{
			"case3",
			"243517698",
			"23568",
		},
		{
			"case4",
			"republican",
			"bcn",
		},
		{
			"case5",
			"democrat",
			"demort",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubsequence2(tt.args); len(got) != len(tt.want) {
				t.Errorf("LongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

var t string

// BenchmarkLongestSubsequence-8   	 2693246	       425 ns/op	     176 B/op	       8 allocs/op
// BenchmarkLongestSubsequence-8   	 5814994	       203 ns/op	     176 B/op	       4 allocs/op
func BenchmarkLongestSubsequence(b *testing.B) {
	s := "243517698"
	var result string
	for i := 0; i < b.N; i++ {
		result = LongestSubsequence(s)
	}
	t = result
}
