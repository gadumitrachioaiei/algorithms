package intervalselection

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestIntervalSelection(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{
			"case1",
			[][]int{{1, 2}, {2, 3}, {2, 4}},
			2,
		},
		{
			"case2",
			[][]int{{1, 5}, {1, 5}, {1, 5}},
			2,
		},
		{
			"case3",
			[][]int{{1, 10}, {1, 3}, {4, 6}, {7, 10}},
			4,
		},
		{
			"case4",
			[][]int{{1, 10}, {1, 3}, {3, 6}, {7, 10}},
			3,
		},
		{
			"case5",
			[][]int{{1, 10}, {1, 6}, {2, 3}, {4, 5}, {7, 10}},
			4,
		},
		{
			"case6",
			[][]int{{1, 3}, {4, 4}, {2, 4}, {1, 3}, {1, 4}, {1, 3}, {1, 1}, {2, 2}, {4, 4}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntervalSelection(tt.args); got != tt.want {
				t.Errorf("IntervalSelection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalSelectionFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	var intervals [][]int
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		a, err := strconv.Atoi(fields[0])
		if err != nil {
			t.Fatal(err)
		}
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			t.Fatal(err)
		}
		intervals = append(intervals, []int{a, b})
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
	if r := IntervalSelection(intervals); r != result {
		t.Fatalf("got result: %d, expected: %d", r, result)
	}
}
