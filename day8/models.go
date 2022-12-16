package day8

import "strconv"

// tree represents a single tree in a 2d map.
type tree struct {
	x      int
	y      int
	height int
}

// coords will return the x and y values of a tree as a string. i.e. x=4, y=3 -> "43"
func (t *tree) coords() string {
	coords := strconv.Itoa(t.x) + ", " + strconv.Itoa(t.y)
	return coords
}
