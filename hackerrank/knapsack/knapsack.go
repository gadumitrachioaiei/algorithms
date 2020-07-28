// Package knapsack solves the unbounded knapsack problem.
// Problem: given an array of integers and a target sum, make up the closest sum of array elements, closer to the target.
// You can use any array element as many times as you want.
//
// Solution:
// Recursive solution based on whether or not an element will be present in the target sum.
// It also has a cache.
package knapsack

func UnboundedKnapsack(k int, arr []int, start int) int {
	if start == len(arr) {
		return 0
	}
	if k < arr[start] {
		return cacheGetSet(k, arr, start+1)
	}
	var r int
	for i := 0; i < k+1; i = i + arr[start] {
		t := i + cacheGetSet(k-i, arr, start+1)
		if r < t {
			r = t
		}
	}
	return r
}

func cacheGetSet(k int, arr []int, i int) int {
	t, ok := cache.Get(k, i)
	if !ok {
		t = UnboundedKnapsack(k, arr, i)
		cache.Set(k, i, t)
	}
	return t
}

var cache = NewCache()

type Element struct {
	k     int
	index int
}

type Cache struct {
	m    map[Element]int
	hits int
}

func NewCache() *Cache {
	return &Cache{m: make(map[Element]int)}
}

func (c *Cache) Set(k int, i int, v int) {
	c.m[Element{k, i}] = v
}

func (c *Cache) Get(k int, i int) (int, bool) {
	v, ok := c.m[Element{k, i}]
	if ok {
		c.hits++
	}
	return v, ok
}
