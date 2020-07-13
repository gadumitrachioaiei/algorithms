// Package queen solves the N-queen attack problem, using backtracking
package queen

// Cohabit determines if it is possible to place n queens on a chess board such that any two don't attack each other
func Cohabit(n int) bool {
	var coords []Coord
	row, col := 0, -1
	for {
		if len(coords) == n {
			break
		}
		found := false
		for j := col + 1; j < n; j++ {
			// search in the last saved row
			if !isAttacked(coords, Coord{row, j}) {
				coords = append(coords, Coord{row, j})
				found = true
				break
			}
		}
		if found {
			// we start our search from after the last coord
			lastCoord := coords[len(coords)-1]
			row = lastCoord.x
			col = -1
			continue
		}
		for i := row + 1; i < n; i++ {
			for j := 0; j < n; j++ {
				if !isAttacked(coords, Coord{i, j}) {
					coords = append(coords, Coord{i, j})
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if len(coords) == 0 {
			break
		}
		if !found {
			// we have to back up
			lastCoord := coords[len(coords)-1]
			coords = coords[:len(coords)-1]
			// but we also have to make sure we start our search from after the last coord
			row = lastCoord.x
			col = lastCoord.y
		}
	}
	return len(coords) == n
}

// Coord is a 2D coordinate
type Coord struct {
	x, y int
}

// isAttacked returns false if q can be attacked by a quuen placed on any of the coordinates
func isAttacked(coords []Coord, q Coord) bool {
	for _, c := range coords {
		if c.x == q.x || c.y == q.y {
			return true
		}
		if (c.x-q.x == c.y-q.y) || (c.x-q.x == q.y-c.y) {
			return true
		}
	}
	return false
}
