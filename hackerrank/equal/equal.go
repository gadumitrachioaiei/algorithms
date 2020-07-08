// Package equal solves the following problem:
// given a list of integers, you can transform it by adding 1,2 or 5 to all but one element
// the goal is for all elements to become equal, and do so with a minimum number of transformations
// e.g.: [1,2,5] ->(2) [3,4,5] ->(2) [5,6,5] ->(1) [6,6,6]
//
// Our solution is like this:
// 1. we reduce the distance between any two elements to less than 5
// 2. we subtract the smallest element from all elements, so our elements will now be in the interval [0,5]
// 3. we want to achieve equality  step by step:
// 4. strategy called "Next", so we try to make an element one element equal to the smallest bigger element
// 5. strategy called "Max", so we try to make the smallest element equal to the biggest
// 6. we choose potential steps to make, depending on the distance that needs to be reduced to 0
// 7. to be fast, we are making all needed steps at once, whenever possible
package equal

import (
	"sort"
)

// Equal returns the smallest number of transformations needed to solve the problem.
func Equal(arr []int) int {
	var countSteps int
	// we want to make the difference between any elements smaller than 5
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		d := arr[i] - arr[0]
		q := d / 5
		countSteps += q
		arr[i] -= q * 5
	}
	// just for convenience, so we can only deal with numbers between with 0 and 5 inclusive
	smallest := arr[0]
	for i := 0; i < len(arr); i++ {
		arr[i] -= smallest
	}
	// this is so we can set a limit on the maximum of steps being performed
	maxRemainingSteps := equalByNeighbour(arr)
	// we are making a map out of the array, as many elements are just repeated, so we can condense them
	m1 := make(map[int]int)
	for i := range arr {
		m1[arr[i]]++
	}
	m2 := clone(m1)
	// we have two possible strategies for calculating the numer of required steps
	steps := 1<<31 - 1
	if c := equal(m1, Next, maxRemainingSteps); c > -1 && c < steps {
		steps = c
	}
	if c := equal(m2, Max, maxRemainingSteps); c > -1 && c < steps {
		steps = c
	}
	if steps == 1<<31-1 {
		return -1
	}
	return countSteps + steps
}

// Strategy describes how we choose which element remains unchanged on each step
type Strategy int

const (
	// Next means we do not change the next big element after min
	Next Strategy = iota
	// Max means we do not change the max element
	Max
)

// equal tries to solve our problem, but it receives a different kind of input,
// and it returns a number of transformations after which all elements become equal.
// m is a map of elements, and their values represent how many times that element is present
// strategy is the applied strategy before we make a transformation
// remainingStepsCount represents how many steps we can make further before giving up
func equal(m map[int]int, strategy Strategy, remainingStepsCount int) int {
	if len(m) == 1 {
		return 0
	}
	if remainingStepsCount < 1 {
		return -1
	}
	// we calculate two keys, and our goal will be for first one to become equal to the second one
	var first, second int
	if strategy == Next {
		first, second = firstSecond(m)
	} else {
		first, second = firstLast(m)
	}
	var step, stepCount int
	// we are making the difference between them as small as possible ( until it becomes 0 ), in terms of steps
	d := second - first
	// 3 and 4 are special cases, as there are two different ways to decrease the step difference
	if d != 3 && d != 4 {
		step = d
		if d > 5 {
			step = 5
		}
		stepCount := makeStepWithSubtraction(m, step, second)
		remainingStepsCount -= stepCount
		c := equal(m, strategy, remainingStepsCount)
		if c == -1 {
			return -1
		}
		return stepCount + c
	}
	// for d = 3 or 4, we have two options, each consisting of two steps
	// for d = 3, we can make up the step difference as: 2 + 1 or 5 - 2
	// for d = 4, we can make up the step difference as: 2 + 2 or 5 - 1
	minStepCount := 1<<31 - 1 // we store the number of steps, received from recursive calls, and we choose their min
	step = d
	// first option
	m1 := clone(m)
	stepCount = makeStepWithSubtraction(m1, step, second)
	if c := equal(m1, strategy, remainingStepsCount-stepCount); c > -1 {
		minStepCount = c + stepCount
	}
	// second option
	// we add 5 to all but second
	stepCount = 2
	m = makeStep(m, 5, second)
	var m2 map[int]int
	if d == 3 {
		// and add 2 to all but first, which now became first+5
		m2 = makeStep(m, 2, first+5)
	} else if d == 4 {
		// and add 1 to second
		m2 = makeStep(m, 1, first+5)
	}
	// just for convenience, so we can only deal with numbers between with 0 and 5 inclusive
	m2 = subtractSmallest(m2)
	if c := equal(m2, strategy, remainingStepsCount-2); c > -1 && c+stepCount < minStepCount {
		minStepCount = c + stepCount
	}
	if minStepCount == 1<<31-1 {
		return -1
	}

	return minStepCount
}

// equalByNeighbour returns a number of transformations that solves our problem, but not necessarily the minimum one
// it does this by making an element equal with its predecessor, starting from the second one
func equalByNeighbour(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	sort.Ints(arr)
	var countSteps int
	var d int
	for i := 1; i < len(arr); i++ {
		d = d + arr[i] - arr[i-1]
		if d == 0 {
			continue
		}
		countSteps += steps(d)
	}
	return countSteps
}

// makeStepWithSubtraction makes the specified step, without changing the value for keyToSKip
// it makes the step as long as possible, until the unchanged key is removed
// it returns the number of steps actually made
// each step is equivalent with subtracting the step from the unchanged key, so we do that
func makeStepWithSubtraction(m map[int]int, step int, keyToSkip int) int {
	v := m[keyToSkip]
	m[keyToSkip-step] += v
	delete(m, keyToSkip)
	if keyToSkip == 3 || keyToSkip == 4 {
		return v * 2
	}
	return v
}

// makeStep returns a new map, made from the given map, by following the specified step
func makeStep(m map[int]int, step int, keyToSkip int) map[int]int {
	mClone := make(map[int]int, len(m))
	for k, v := range m {
		if k != keyToSkip {
			mClone[k+step] = v
		}
	}
	if m[keyToSkip] > 1 {
		mClone[keyToSkip+step] = m[keyToSkip] - 1
	}
	mClone[keyToSkip] += 1
	return mClone
}

// subtractSmallest returns a new map, by substracting the smallest subkey from all the elements
func subtractSmallest(m map[int]int) map[int]int {
	first := 1<<31 - 1
	for k := range m {
		if k < first {
			first = k
		}
	}
	m2 := make(map[int]int, len(m))
	for k := range m {
		m2[k-first] = m[k]
	}
	return m2
}

// firstSecond returns the first two keys of the map, in ascending order, that are different
func firstSecond(m map[int]int) (int, int) {
	first, second := -1, -1
	for k := range m {
		if first == -1 {
			first, second = k, k
			continue
		}
		if k < first {
			first, second = k, first
		} else if k < second {
			second = k
		} else if first == second {
			second = k
		}
	}
	return first, second
}

// firstLast returns the first two keys of the map, in ascending order, that are different
func firstLast(m map[int]int) (int, int) {
	first, last := -1, -1
	for k := range m {
		if first == -1 {
			first, last = k, k
			continue
		}
		if k < first {
			first = k
		} else if k > last {
			last = k
		}
	}
	return first, last
}

// steps returns the smallest number of 1, 2 and 5 that together add up to d
func steps(d int) int {
	q1 := d / 5
	d = d - q1*5
	q2 := d / 2
	d = d - q2*2
	return q1 + q2 + d
}

// clone returns a new identical map
func clone(m map[int]int) map[int]int {
	n := make(map[int]int)
	for k, v := range m {
		n[k] = v
	}
	return n
}
