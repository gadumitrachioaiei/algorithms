// https://www.hackerrank.com/challenges/stone-division-2/problem
package stonedivision

import (
	"sort"
)

func StoneDivision(n int, s []int) int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	cache = make(map[int]int)
	return stoneDivision2(n, s)
}

func stoneDivision2(n int, s []int) int {
	var factors = make([]int, 0, 10)
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] == n {
			continue
		}
		if n%s[i] != 0 {
			continue
		}
		factor := false
		for j := 0; j < len(factors); j++ {
			if factors[j]%s[i] == 0 {
				factor = true
				break
			}
		}
		if factor {
			continue
		}
		factors = append(factors, s[i])
		v, ok := cache[s[i]]
		if !ok {
			v = stoneDivision2(s[i], s[i+1:])
			cache[s[i]] = v
		}
		if r := (n/s[i])*v + 1; r > result {
			result = r
		}
	}
	return result
}

var cache map[int]int
