package stringpairs

import (
	"strings"
	"testing"
)



func TestSolution(t *testing.T) {
	tcs := []struct{
		input, solution string
	}{
		{"", ""},
		{"ACCAABBC", "AC"},
		{"BABABA", "BABABA"},
		{"ABCBBCBA", ""},
		{"ABCCBAABCCBAABCCBA", ""},
	}
	for _, tc := range tcs {
		if solution := Solution(tc.input); solution != tc.solution {
			t.Errorf("got:\n%s\nexpected:\n%s\n", solution, tc.solution)
		}
	}
}

var solution string

func BenchmarkSolution(b *testing.B) {
	input := "ACCAABBC"
	input = strings.Repeat(input, 10000)
	var s string
	for i := 0; i<b.N; i++ {
		s = Solution(input)
	}
	solution = s
}
