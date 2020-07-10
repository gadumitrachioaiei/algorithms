// Package candies solves this problem:
// given an array of numbers, associate to each of its elements, a number such that:
// 1. it is at least 1
// 2. neighbouring numbers must follow the same ascending or descending relationship, as in the original array
// 3. if two neighbouring elements are equal, there is no restriction between each other, other than the minimal requirement of 1
// 4. the sum of all produced elements must be minimal, respecting the above requirements
package candies

// Candies returns the sum of elements of the new elements
func Candies(a []int) int {
	if len(a) < 2 {
		if a[0] == a[1] {
			return 2
		}
		return 3
	}
	b := make([]int, len(a)+1)
	// establish the initial type of relation
	var rel Rel
	// point of extrem, either a minimum, or a maximum
	var pExtreme int
	if a[0] < a[1] {
		rel = Asc
	} else if a[0] > a[1] {
		rel = Desc
	}
	if rel == Eq {
		b[0], b[1] = 1, 1
		pExtreme = 1
	}
	// we append another equal element, so that our for loops associates elements to all of the original array
	// otherwise, we needed to have written similar code, outside of the for loop, to process the last elements
	a = append(a, a[len(a)-1])
	for i := 2; i < len(a); i++ {
		var nextRel Rel
		d := a[i] - a[i-1]
		switch {
		case d == 0:
			nextRel = Eq
		case d < 0:
			nextRel = Desc
		default:
			nextRel = Asc
		}
		// we handle changes in monotony
		// uf current point becomes a point of minimum
		if (rel == Eq && nextRel == Asc) ||
			(rel == Desc && nextRel == Asc) ||
			(rel == Desc && nextRel == Eq) {
			// we associate numbers for all elements to the left of the minimum, up to the last extreme
			b[i-1] = 1
			for k := i - 1; k > pExtreme; k-- {
				b[k] = i - k
			}
			if b[pExtreme] < i-pExtreme {
				b[pExtreme] = i - pExtreme
			}
			// we are recording a point of minimum
			rel = nextRel
			pExtreme = i - 1
		}
		// if current point is a point of maximum
		if (rel == Eq && nextRel == Desc) ||
			(rel == Asc && nextRel == Desc) ||
			(rel == Asc && nextRel == Eq) {
			// we associate numbers for all elements to the left of the maximum, up to the last extreme
			for k := pExtreme; k < i; k++ {
				b[k] = k - (pExtreme - 1)
			}
			// we are recording a point of maximum
			rel = nextRel
			pExtreme = i - 1
		}
		// all intermediary elements that lie on an eq rel can be set as 1
		if nextRel == Eq {
			b[i] = 1
			pExtreme = i
		}
	}
	var sum int
	// we stop before the last element, because it was artifically added by us
	for i := 0; i < len(b)-1; i++ {
		sum += b[i]
	}
	return sum
}

// Rel categorises the monontony of the array so far
type Rel int

const (
	// last elements are equal
	Eq Rel = iota
	// last elements are in ascending order
	Asc
	// last elements are in descending order
	Desc
)
