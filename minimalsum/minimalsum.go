// Package minimalsum solves the problem:
// What is the minimal sum that we can obtain from 3 numbers,
// after we perform any number of the following operations:
// 1. you can only subtract 1 or 2, from any number
// 2. you can not subtract from the same number twice in a row
// 3. you can not subtract past 0

// Another variant of the problem, from codility:
// You are given  a count of occurences of each of the letters, "a", "b", "c", as integers: A, B, C.
// Find out the length of the longest string we can make with them, considering we cannot have the same letter 3 consecutive times.
// Example:
// A = 6, B = 1, C = 1, we can make up the string: aabaacaa

// Solution:
// The heuristic is to always make sure that we can continue with another operation after the current one.
// We can make sure that we always subtract from the first two biggest occurrences, such that the difference between them stays as small as possible.
package minimalsum

import (
	"sort"
)

// MinimalSum returns the minimal sum that remains from the 3 numbers after the allowed operations.
func MinimalSum(A, B, C int) int {
	l := []Element{{A, 0}, {B, 1}, {C, 2}}
	canNotChoose := 3
	for {
		sort.Slice(l, func(i, j int) bool { return l[i].count < l[j].count })
		if (l[2].index != canNotChoose && l[2].count == 0) ||
			(l[2].index == canNotChoose && l[1].count == 0) {
			break
		}
		if l[2].index != canNotChoose {
			min := 2
			if l[2].count < min {
				min = l[2].count
			}
			l[2].count -= min
			canNotChoose = l[2].index
		} else {
			l[1].count -= 1
			canNotChoose = l[1].index
		}
	}
	return l[0].count + l[1].count + l[2].count
}

type Element struct {
	count int
	index int
}
