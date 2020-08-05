// Package steadygene: https://www.hackerrank.com/challenges/bear-and-steady-gene/problem
// Problem:
//
// We are given a string that contains at most four different characters, of length n ( multiple of 4 ).
// We need to find the smallest substring, that we can replace with any other string of same size,
// such that all four characters will be found in the original string in equal number, n / 4
// Solution:
//
// We find the minimal length of a substring that needs to be replaced.
// We then loop through all substrings of greater length and check that it contains at least the crachters that need to be replaced.
// We go to the next substring by advancing its starting char or its size as much as possible.
// And we make optimizations about caching previous calculations of contained chars.
package steadygene

func SteadyGene(gene string) int {
	// map characters to their number of occurrences
	n := len(gene)
	t := n / 4
	arr := make([]int, 100)
	for i := range gene {
		arr[gene[i]]++
	}
	// find the characters that are over the limit t and the indexes of those chars in tha over array
	// also set the count of characters that are over the limit
	length := 0
	over := make([]int, 100)
	indexesOver := make([]int, 0, 3)
	for i, v := range arr {
		diff := v - t
		if diff > 0 {
			length += diff
			over[i] = diff
			indexesOver = append(indexesOver, i)
		}
	}
	// we cache the count array for the substring of current size and starting at char 0
	cacheForCurrentSize := make([]int, 100)
	for i := 0; i < length; i++ {
		cacheForCurrentSize[gene[i]]++
	}
	// arr1 is the array with counts for current size, and varying starting indexes
	// it is initialized with the counts for index 0
	arr1 := append(cacheForCurrentSize[:0:0], cacheForCurrentSize...)
	size := length // the current size of analyzed substring
	// how much do we need to advance the current size so we have a chance of finding a steady gene substring
	index0Diff := 1<<31 - 1
	// we start looking at substrings of length equal to size
	for size < n {
		found := false
		j := 0
		// we start looking at substrings starting at index j
		for j < n-size+1 {
			d := diff(arr1, over, indexesOver)
			if d == 0 {
				found = true
				break
			}
			// on the next size, we can advance with this many chars
			if index0Diff > d {
				index0Diff = d
			}
			if j+d > n-size {
				break
			}
			// we advance the current starting index with d
			for k := j; k < j+d; k++ {
				arr1[gene[k]]--
				arr1[gene[k+size]]++
			}
			j += d
		}
		if found {
			break
		}
		if size+index0Diff > n-1 {
			break
		}
		// we reinitialize the array with counts and advance the size with index0Diff
		arr1 = append(cacheForCurrentSize[:0:0], cacheForCurrentSize...)
		for k := size; k < size+index0Diff; k++ {
			arr1[gene[k]]++
		}
		cacheForCurrentSize = append(arr1[:0:0], arr1...)
		size += index0Diff
	}
	return size
}

// diff returns the different in counts between arr2 and arr1, over the indexes in indexesOver
func diff(arr1, arr2 []int, indexesOver []int) int {
	over := 0
	for _, index := range indexesOver {
		if d := arr2[index] - arr1[index]; d > 0 {
			over += d
		}
	}
	return over
}
