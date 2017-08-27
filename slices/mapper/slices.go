package mapper

type Indexer interface {
	Len() int
}

type Equaler interface {
	Equal(i, j int) bool
}

type IndexEqualer interface {
	Indexer
	Equaler
}

// Map maps elements in slice a to elements in slice b, it assumes a and b are sorted such that
// an element in a is mapped to indexes greater than indexes previously mapped
// If two adjacent elements in a are equal, they are mapped to same elements in b
// an element in a is not required to be mapped
// MapData from test file is an example of a caller
func Map(a IndexEqualer, b Indexer, isMap func(i, j int) bool, fMap func(i, j, k int)) {
	startB := 0
	startBLoop := 0
	idxLeft, idxRight := 0, 0 // indexes in b that previous element in a maps to
	for i := 0; i < a.Len(); i++ {
		if i > 0 && a.Equal(i-1, i) {
			fMap(i, idxLeft, idxRight)
			continue
		}
		idxLeft, idxRight = 0, 0
		found := false
		for j := startBLoop; j < b.Len(); j++ {
			mapped := isMap(i, j)
			if mapped && !found {
				startB = j
				found = true
			} else if !mapped && found {
				fMap(i, startB, j)
				idxLeft, idxRight = startB, j
				startBLoop = j
				found = false
				break
			}
		}
		if found {
			fMap(i, startB, b.Len())
			idxLeft, idxRight = startB, b.Len()
		}
	}
}
