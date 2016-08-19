package transformer

import (
	"reflect"
	"testing"
)

func TestGetLongestChain(t *testing.T) {
	stop := map[int]struct{}{1: struct{}{}}
	limit := 100000
	maxChainLength := -1
	expectedResult := -1
	for n := 1; n < limit+1; n++ {
		chain := findChain(n, stop)
		if len(chain) > maxChainLength {
			maxChainLength = len(chain)
			expectedResult = n
		}
	}
	alg := &Algorythm{}
	result := alg.GetLongestChain(limit)
	if expectedResult != result {
		t.Fatalf("The longest chain is for number: %d but we calculate it to be for number: %d", expectedResult, result)
	}
}

type TestGolden struct {
	n, m int
	path []int
}

var tests = []TestGolden{
	{10, 10, []int{10}},
	{10, 11, []int{10, 20, 40, 13, 26, 52, 17, 34, 11}},
	{10, 12, []int{10, 3, 6, 12}},
	{10, 100, []int{10, 20, 40, 13, 26, 52, 17, 34, 11, 22, 44, 88, 29, 58, 19, 38, 76, 25, 50, 100}},
	{7, 9, []int{7, 14, 28, 9}},
}

func TestGetShortestPath(t *testing.T) {
	for _, test := range tests {
		path := GetShortestPath(test.n, test.m)
		if !reflect.DeepEqual(test.path, path) {
			t.Logf("Expected path between %d and %d to be: %v but got: %v", test.n, test.m, test.path, path)
			t.Fail()
		}
	}
}
