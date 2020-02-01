package coveringsubstring

import (
	"fmt"
	"strings"
	"testing"
)

func TestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "1",
			args: "A",
			want: "A",
		},
		{
			name: "2",
			args: "AA",
			want: "A",
		},
		{
			name: "3",
			args: "ABCD",
			want: "ABCD",
		},
		{
			name: "4",
			args: "ABCBCD",
			want: "ABCBCD",
		},
		{
			name: "5",
			args: "AAABBBCDDD",
			want: "ABBBCD",
		},
		{
			name: "6",
			args: "ABCBDEEDF",
			want: "ABCBDEEDF",
		},
		{
			name: "7",
			args: "ABCBAEEDF",
			want: "CBAEEDF",
		},
		{
			name: "8",
			args: "ABCCBA",
			want: "ABC",
		},
		{
			name: "9",
			args: "ABCABC",
			want: "ABC",
		},
		{
			name: "10",
			args: "\u07ff\u07e6\u07e6\u07ff",
			want: "\u07ff\u07e6",
		},
		{
			name: "11",
			args: "r\U0003ff4a",
			want: "r\U0003ff4a",
		},
		{
			name: "12",
			args: "0000000000000000\ufffd",
			want: "0\ufffd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substring(tt.args); got != tt.want {
				t.Errorf("Substring() = %+q, want %+q", got, tt.want)
			}
		})
	}
}

var ss string

func BenchmarkSubstring(b *testing.B) {
	run := func(b *testing.B, s string) {
		for i := 0; i < b.N; i++ {
			ss = Substring(s)
		}
	}
	runR := func(b *testing.B, s string) {
		for i := 0; i < b.N; i++ {
			ss = SubstringR(s)
		}
	}
	inputs := []string{
		"ABCDEABF",
		"ABCDEABCEABCDFGHIJKLABCLMNOPQRSTUVTUVXYZ",
		"ABCDEFGHIJKLMNOPQRSTUVXYZ",
		strings.Repeat("ABCDEABF", 100),
		strings.Repeat("ABCDEABCEABCD", 100),
	}
	b.ResetTimer()
	for i, f := range []func(b *testing.B, s string){run, runR} {
		for j, s := range inputs {
			b.Run(fmt.Sprintf("%d_%d", i, j), func(b *testing.B) {
				f(b, s)
			})
		}
	}
}
