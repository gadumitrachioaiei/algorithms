package knapsack

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test_unboundedKnapsack(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{12, []int{1, 6, 9}},
			12,
		},
		{
			"case2",
			args{9, []int{3, 4, 4, 4, 8}},
			9,
		},
		{
			"case3",
			args{10, []int{3, 2, 4}},
			10,
		},
		{
			"case4",
			args{13, []int{3, 7, 9, 11}},
			13,
		},
		{
			"case5",
			args{11, []int{3, 7, 9}},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache = NewCache()
			if got := UnboundedKnapsack(tt.args.k, tt.args.arr, 0); got != tt.want {
				t.Errorf("UnboundedKnapsack() = %v, want %v", got, tt.want)
			}
			t.Log("hits", cache.hits)
		})
	}
}

func TestUnboundedKnapsackFromInput(t *testing.T) {
	inputPath := "input.txt"
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		t.Fatal(err)
	}
	fields := strings.Fields(string(input))
	k, err := strconv.Atoi(fields[0])
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
	if r := UnboundedKnapsack(k, arr, 0); r != result {
		t.Fatalf("got result: %d, expected: %d", r, result)
	}
}
