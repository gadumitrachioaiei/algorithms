// Package largestrectangle : https://www.hackerrank.com/challenges/largest-rectangle/problem
// Problem:
//
// You are given an array of heights of buildings. Find the largest rectangular area that can be constructed with a continuous list of buildings,
// knowing that the area is the product between the minimum height of the buildings and their count.
package largestrectangle

// LargestRectangle returns the are of the largest continuous rectangle, formed with continuous heights
func LargestRectangle(h []int) int {
	m := closestSmallerNeighbours(h)
	area := 0
	for i := 0; i < len(m); i++ {
		if a := h[i] * (m[i][1] - m[i][0] - 1); a > area {
			area = a
		}
	}
	return area
}

// closestSmallerNeighbours maps each index in h with indexes of the closest neighbours
// to the left and right, whose values are smaller than the index value.
func closestSmallerNeighbours(h []int) [][2]int {
	m := make([][2]int, len(h))
	i := 0
	for i < len(h) {
		leftIndex, rightIndex := closestSmallerNeighboursForIndex(h, i)
		m[i] = [2]int{leftIndex, rightIndex}
		for l := i + 1; l < rightIndex; l++ {
			m[l][1] = rightIndex
			for n := l + 1; n < rightIndex; n++ {
				if h[n] < h[l] {
					m[l][1] = n
					break
				}
			}
			m[l][0] = leftIndex
			for n := l - 1; n > leftIndex; n-- {
				if h[n] < h[l] {
					m[l][0] = n
					break
				}
			}
		}
		i = rightIndex
	}
	return m
}

// closestSmallerNeighboursForIndex is like closestSmallerNeighbours, but for a given index.
func closestSmallerNeighboursForIndex(h []int, i int) (leftIndex, rightIndex int) {
	leftIndex, rightIndex = -1, len(h)
	for k := 0; k < i; k++ {
		if h[k] < h[i] && leftIndex < k {
			leftIndex = k
		}
	}
	for j := i + 1; j < len(h); j++ {
		if h[j] < h[i] {
			rightIndex = j
			break
		}
	}
	return leftIndex, rightIndex
}
