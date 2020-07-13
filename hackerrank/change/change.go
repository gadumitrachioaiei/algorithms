// Package change solves the change problem:
// Given some type of coins, in how many ways can you return a given change ?
// it is classic dynamic programming, i.e: recursive, overlapping sub cases with cache
package change

import (
	"fmt"
	"sort"
)

// Change solves the change problem
func Change(units int, coins []int) int {
	sort.Ints(coins)
	return change(units, coins)
}

func change(units int, coins []int) int {
	if units == 0 {
		return 1
	}
	var result int
	index := sort.SearchInts(coins, units)
	if index < len(coins) && coins[index] == units {
		result++
	}
	coins = coins[:index]
	for i, c := range coins {
		for j := c; j < units+1; j = j + c {
			r, ok := cache.Get(units-j, coins[i+1:])
			if !ok {
				r = change(units-j, coins[i+1:])
				cache.Set(units-j, coins[i+1:], r)
			}
			result += r
		}
	}
	return result
}

type Element struct {
	coins string
	units int
}

type Cache struct {
	data map[Element]int
}

func (c Cache) Get(units int, a []int) (int, bool) {
	r, ok := c.data[Element{fmt.Sprint(a), units}]
	return r, ok
}

func (c Cache) Set(units int, a []int, result int) {
	c.data[Element{fmt.Sprint(a), units}] = result
}

var cache = Cache{data: make(map[Element]int)}
