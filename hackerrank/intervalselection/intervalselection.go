// Package intervalselection : https://www.hackerrank.com/challenges/interval-selection/problem
// Solution:
//
// Have all ends of the intervals in a new sorted array.
// Loop through it, and when you find 3 open intervals, remove the one that stretches farthest.
package intervalselection

import (
	"sort"
)

func IntervalSelection(intervals [][]int) int {
	ends := make([]End, 2*len(intervals))
	for i := 0; i < len(intervals); i++ {
		ends[2*i] = End{intervals[i][0], i, true}
		ends[2*i+1] = End{intervals[i][1], i, false}
	}
	sort.Slice(ends, func(i, j int) bool {
		if ends[i].v == ends[j].v {
			if ends[i].open {
				// we would like open intervals to be set before closed ones
				return true
			}
			return false
		}
		return ends[i].v < ends[j].v
	})
	var dismissed int
	open := make(map[int]bool)
	for i := 0; i < len(ends); i++ {
		if !ends[i].open {
			delete(open, ends[i].index)
			continue
		}
		open[ends[i].index] = true
		if len(open) == 3 {
			removeMaximum(open, intervals)
			dismissed++
		}
	}
	return len(intervals) - dismissed
}

type End struct {
	v     int
	index int
	open  bool
}

func removeMaximum(m map[int]bool, intervals [][]int) {
	type Size struct {
		v     int
		index int
	}
	sizes := make([]Size, 0, 3)
	for i := range m {
		size := intervals[i][1]
		sizes = append(sizes, Size{size, i})
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i].v < sizes[j].v
	})
	delete(m, sizes[2].index)
}
