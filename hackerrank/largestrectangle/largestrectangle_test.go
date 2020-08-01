package largestrectangle

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestLargestRectangle(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"case1",
			[]int{3, 2, 3},
			6,
		},
		{
			"case2",
			[]int{1, 2, 3, 4, 5},
			9,
		},
		{
			"case3",
			[]int{1, 3, 5, 9, 11},
			18,
		},
		{
			"case4",
			[]int{11, 11, 10, 10, 10},
			50,
		},
		{
			"case5",
			[]int{6320, 6020, 6098, 1332, 7263, 672, 9472, 2838, 3401, 9494},
			18060,
		},
		{
			"case6",
			[]int{6873, 7005, 1372, 5438, 1323, 9238, 9184, 2396, 4605, 162},
			18368,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangle(tt.args); got != tt.want {
				t.Errorf("LargestRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

//             6873   7005,  1372, 5438,  1323,  9238, 9184, 2396, 4605, 162
// correct:   [[-1 2] [0 2] [-1 4] [2 4] [-1 9] [4 6] [4 7] [4 9] [7 9] [-1 10]]
// incorrect: [[-1 2] [0 2] [-1 4] [2 4] [-1 9] [4 9] [4 9] [4 9] [7 9] [-1 10]]

func TestLargestRectangleFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	var h []int
	for _, field := range strings.Fields(string(input)) {
		height, err := strconv.Atoi(field)
		if err != nil {
			t.Fatal(err)
		}
		h = append(h, height)
	}
	outputPath := "output.txt"
	output, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	result, err := strconv.Atoi(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if got := LargestRectangle(h); got != result {
		t.Errorf("LargestRectangle() = %v, want %v", got, result)

	}
}
