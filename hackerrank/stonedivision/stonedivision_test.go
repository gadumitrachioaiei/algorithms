package stonedivision

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestStoneDivision(t *testing.T) {
	type args struct {
		n int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{12, []int{2, 3, 4}},
			4,
		},
		{
			"case2",
			args{64, []int{2, 4, 8, 16, 64}},
			29,
		},
		{
			"case3",
			args{1, []int{1, 2}},
			0,
		},
		{
			"case4",
			args{6, []int{3}},
			1,
		},
		{
			"case5",
			args{64, []int{2, 4, 8, 16, 32, 64}},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StoneDivision(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("StoneDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoneDivisionFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	fields := strings.Fields(string(input))
	n, err := strconv.Atoi(fields[0])
	if err != nil {
		t.Fatal(err)
	}
	var arr []int
	for _, field := range fields[1:] {
		el, err := strconv.Atoi(field)
		if err != nil {
			t.Fatal(err)
		}
		arr = append(arr, el)
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
	if r := StoneDivision(n, arr); r != result {
		t.Fatalf("got result: %d, expected: %d", r, result)
	}
}
