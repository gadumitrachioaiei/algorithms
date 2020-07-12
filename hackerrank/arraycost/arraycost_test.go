package arraycost

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestCost(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"case1",
			[]int{5, 3, 5, 5, 1},
			12,
		}, {
			"case2",
			[]int{6, 4, 6, 6, 1},
			16,
		}, {
			"case3",
			[]int{10, 1, 10, 1, 10},
			36,
		}, {
			"case4",
			[]int{9, 9, 9, 9},
			24,
		}, {
			"case5",
			[]int{5, 1, 5, 5, 1, 5},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cost(tt.args); got != tt.want {
				t.Errorf("Cost() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestEqualFromInput tests Equal using some big inputs
func TestEqualFromInput(t *testing.T) {
	inputPath := "input.txt"
	f, err := os.Open(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	outputPath := "output.txt"
	output, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	var outputSteps []int
	for _, el := range bytes.Fields(output) {
		el, err := strconv.Atoi(string(el))
		if err != nil {
			t.Fatal(err)
		}
		outputSteps = append(outputSteps, el)
	}
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 64*1024*1024)
	scanner.Buffer(buf, 64*1024*1024)
	tests := -1
	var failedTests []int
	for scanner.Scan() {
		tests++
		var arr []int
		lastIndex := 0
		data := scanner.Bytes()
		for i, c := range data {
			var index int
			if c == ' ' {
				index = i
			} else if i == len(data)-1 {
				index = i + 1
			} else {
				continue
			}
			el, err := strconv.Atoi(string(data[lastIndex:index]))
			if err != nil {
				t.Fatal(err)
			}
			lastIndex = i + 1
			arr = append(arr, el)
		}
		if r := Cost(arr); r != outputSteps[tests] {
			failedTests = append(failedTests, tests)
			t.Errorf("calculated: %d, expected: %d", r, outputSteps[tests])
		}
	}
	t.Log("failed tests", failedTests)
}
