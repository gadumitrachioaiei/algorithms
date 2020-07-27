package substringdiff

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestSubstringDiff(t *testing.T) {
	type args struct {
		k  int
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
			args{1, "abcd", "bbca"},
			3,
		},
		{
			"case2",
			args{2, "tabriz", "torino"},
			4,
		},
		{
			"case3",
			args{0, "abacba", "xbcaba"},
			3,
		},
		{
			"case4",
			args{3, "helloworld", "yellomarin"},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringDiff(tt.args.k, tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("SubstringDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSubstringDiffFromInput tests SubstringDiff using some big inputs
func TestSubstringDiffFromInput(t *testing.T) {
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
	var results []int
	for _, el := range bytes.Fields(output) {
		el, err := strconv.Atoi(string(el))
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, el)
	}
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 64*1024*1024)
	scanner.Buffer(buf, 64*1024*1024)
	tests := -1
	var failedTests []int
	for scanner.Scan() {
		tests++
		data := scanner.Bytes()
		fields := strings.Fields(string(data))
		if len(fields) != 3 {
			t.Fatalf("got %d spaces, expected: 2", len(fields))
		}
		k, err := strconv.Atoi(fields[0])
		if err != nil {
			t.Fatalf("got data: %s, expected int", fields[0])
		}
		s1, s2 := fields[1], fields[2]
		now := time.Now()
		if r := SubstringDiff(k, s1, s2); r != results[tests] {
			failedTests = append(failedTests, tests)
			t.Fatalf("test: %d, got result: %d, expected: %d", tests, r, results[tests])
		}
		t.Log("test", tests, time.Since(now))
	}
}
