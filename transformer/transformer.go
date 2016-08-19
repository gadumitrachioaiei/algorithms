// Package transformer assumes that a chain of transformations stops at a number smaller than max int64
package transformer

import "log"

// Algorythm implements longest chain of tranformations up to a number
type Algorythm struct {
	n               int         // we want to calculate longest chain for all numbers less or equal with n
	chainLengths    []int       // slice with chain lengths for all numbers less or equal with n
	chainLengthsMap map[int]int // map with chain lengths for number strictly greater than n, hopefully not very big so we save memory
}

// what is the length of chain of transformations from number to 1
func (alg *Algorythm) find(number int) int {
	if number == 1 {
		// nothing to do
		return 0
	}
	if number < alg.n+1 {
		// we look for the chain length in the slice
		if alg.chainLengths[number] != 0 {
			// if already calculated
			return alg.chainLengths[number]
		}
	} else if alg.chainLengthsMap[number] != 0 {
		// if already calculated
		return alg.chainLengthsMap[number]
	}
	// claculate chainLength for a new number
	chainLength := 0
	if number%2 == 0 {
		chainLength = alg.find(number/2) + 1
	} else {
		chainLength = alg.find(3*number+1) + 1
	}
	// savw chain length in the slice if number smaller than n else in the map
	if number < alg.n+1 {
		alg.chainLengths[number] = chainLength
	} else {
		alg.chainLengthsMap[number] = chainLength
	}
	return chainLength
}

// GetLongestChain returns the smallest number up to n that produces the longest chain
func (alg *Algorythm) GetLongestChain(n int) int {
	if n < 1 {
		log.Fatalf("n should be bigger than 0")
	}
	alg.n = n
	alg.chainLengths = make([]int, n+1)
	alg.chainLengthsMap = make(map[int]int)
	result := -1
	maxChainLength := -1
	for i := 1; i < n+1; i++ {
		chainLength := alg.find(i)
		if chainLength > maxChainLength {
			maxChainLength = chainLength
			result = i
		}
	}
	return result
}

// findChain finds the chain of transformations for the number, up to the first number found in stop map
func findChain(number int, stop map[int]struct{}) []int {
	chain := []int{number}
	for {
		if _, ok := stop[number]; ok {
			break
		}
		if number%2 == 0 {
			number /= 2
		} else {
			number = 3*number + 1
		}
		chain = append(chain, number)
	}
	return chain
}

// GetShortestPath returns shortest path between the numbers walking down
// from the first number to the common chain element and back up to the second element
func GetShortestPath(n, m int) []int {
	// find the chain for n
	stop := map[int]struct{}{1: struct{}{}}
	nChain := findChain(n, stop)
	stop = make(map[int]struct{})
	for _, chainElement := range nChain {
		stop[chainElement] = struct{}{}
	}
	mChain := findChain(m, stop)
	// join the two chains
	commonChainElement := mChain[len(mChain)-1]
	var result []int
	for i, chainElement := range nChain {
		if chainElement == commonChainElement {
			result = nChain[:i]
			break
		}
	}
	for i := len(mChain) - 1; i > -1; i-- {
		result = append(result, mChain[i])
	}
	return result
}
