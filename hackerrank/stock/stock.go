package stock

import (
	"sort"
)

func Stockmax(prices []int) int {
	elements := make([]Element, len(prices))
	for i := 0; i < len(prices); i++ {
		elements[i] = Element{prices[i], i}
	}
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].v == elements[j].v {
			return elements[i].i > elements[j].i
		}
		return elements[i].v > elements[j].v
	})
	var result int
	lastIndex := -1
	for i := 0; i < len(elements); i++ {
		for j := lastIndex + 1; j < elements[i].i; j++ {
			result += elements[i].v - prices[j]
		}
		if elements[i].i > lastIndex {
			lastIndex = elements[i].i
		}
	}
	return result
}

type Element struct {
	v int // value
	i int // index
}
