package coveringsubstring

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"

	fuzz "github.com/google/gofuzz"
)

// TestSubstring tests Substring by comparing its result with the test substringTest
// it uses random generation of utf-8 strings
func TestSubstring(t *testing.T) {
	f := fuzz.New()
	var s string
	for i := 0; i < 100; i++ {
		f.Fuzz(&s)
		expected := substringTest(s)
		calculated := Substring(s)
		if len(calculated) != len(expected) {
			t.Fatalf("not the same length, s:\n%s\ncalculated:\n%s\nexample expected:\n%s\n", s, calculated, expected)
		}
		if !reflect.DeepEqual(unique(calculated), unique(expected)) {
			t.Fatalf("not the same unique character, s:\n%s\ncalculated:\n%s\nexample expected:\n%s\n", s, calculated, expected)
		}
	}
}

func TestSubstringExpected(t *testing.T) {
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
		{
			name: "13",
			args: "e腭3ʋ圔ǋ:¹鼉ĖY{{",
			want: "e腭3ʋ圔ǋ:¹鼉ĖY{",
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

// substring calculates the covering substring, so as the code as more readable without carrying about performance
// because of performance not suitable for long strings
func substringTest(s string) string {
	if !utf8.Valid([]byte(s)) {
		panic("not utf8: " + s)
	}
	result := s
	ssLength := len([]rune(result))
	c := len(unique(result))
	for i := range s {
		var size int
		for j := i; j < len(s); j += size {
			_, size = utf8.DecodeRuneInString(s[j:])
			if m := unique(s[i : j+size]); len(m) == c {
				if l := len([]rune(s[i : j+size])); l < ssLength {
					result = s[i : j+size]
					ssLength = l
				}
			}
		}
	}
	return result
}
