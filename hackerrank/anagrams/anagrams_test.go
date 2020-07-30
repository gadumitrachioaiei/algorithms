package anagrams

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestSherlockAndAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			"case1",
			"mom",
			2,
		},
		{
			"case2",
			"abba",
			4,
		},
		{
			"case3",
			"abcd",
			0,
		},
		{
			"case4",
			"ifailuhkqq",
			3,
		},
		{
			"case5",
			"kkkk",
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagrams(tt.args); got != tt.want {
				t.Errorf("Anagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSherlockAndAnagramsFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	ss := strings.Fields(string(input))
	outputPath := "output.txt"
	output, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	results := strings.Fields(string(output))
	for i, s := range ss {
		result, err := strconv.Atoi(results[i])
		if err != nil {
			t.Fatal(err)
		}
		if r := Anagrams(s); r != result {
			t.Fatalf("testcase %d got result: %d, expected: %d", i, r, result)
		}
	}
}
