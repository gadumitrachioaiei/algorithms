package commonchild

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestCommonChild(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{"ABCD", "ABDC"},
			3,
		},
		{
			"case2",
			args{"HARRY", "SALLY"},
			2,
		},
		{
			"case3",
			args{"AA", "BB"},
			0,
		},
		{
			"case4",
			args{"SHINCHAN", "NOHARAAA"},
			3,
		},
		{
			"case5",
			args{"ABCDEF", "FBDAMN"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonChild(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CommonChild() = %v, want %v", got, tt.want)
			}
			t.Log("hits", cache.hits)
		})
	}
}

func TestCommonChildFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	fields := strings.Fields(string(input))
	outputPath := "output.txt"
	output, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	result, err := strconv.Atoi(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if r := CommonChild(fields[0], fields[1]); r != result {
		t.Fatalf("got result: %d, expected: %d", r, result)
	}
	t.Log("hits", cache.hits)
}
