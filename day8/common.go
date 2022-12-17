package day8

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
)

type viewAngle int32

// viewAngle is the different angles that you can try to view trees from
const (
	top viewAngle = iota
	bottom
	left
	right
)

// buildTreePlane will read the input and build the 2D array of treeHeight values
func buildTreePlane() (plane [][]treeHeight) {
	for line := range common.ReadLinesFromFile("day8") {
		var row []treeHeight

		for _, rawHeight := range line {
			height, err := strconv.Atoi(string(rawHeight))
			if err != nil {
				panic(err)
			}

			row = append(row, treeHeight(height))
		}
		plane = append(plane, row)
	}
	return plane
}

// isVisible will check if a tree (by x, y) is visible in a plane from a certain angle
func isVisible(x int, y int, height treeHeight, plane [][]treeHeight, angle viewAngle) (visible bool) {
	visible = true

	switch angle {
	case top:
		for iy := y - 1; iy >= 0; iy-- {
			if plane[iy][x] >= height {
				visible = false
				break
			}
		}
	case bottom:
		for iy := y + 1; iy < len(plane); iy++ {
			if plane[iy][x] >= height {
				visible = false
				break
			}
		}
	case left:
		for ix := x - 1; ix >= 0; ix-- {
			if plane[y][ix] >= height {
				visible = false
				break
			}
		}
	case right:
		for ix := x + 1; ix < len(plane[0]); ix++ {
			if plane[y][ix] >= height {
				visible = false
				break
			}
		}
	default:
		panic(fmt.Sprintf("Unrecognized angle: %d", angle))
	}
	return visible
}

func getScenicScore(x int, y int, height treeHeight, plane [][]treeHeight) (scenicScore int) {
	var topScore int
	var bottomScore int
	var leftScore int
	var rightScore int

	// Top iterator
	for iy := y - 1; iy >= 0; iy-- {
		topScore = y - iy
		if plane[iy][x] >= height {
			break
		}
	}

	// Bottom iterator
	for iy := y + 1; iy < len(plane); iy++ {
		bottomScore = iy - y
		if plane[iy][x] >= height {
			break
		}
	}

	// Left iterator
	for ix := x - 1; ix >= 0; ix-- {
		leftScore = x - ix
		if plane[y][ix] >= height {
			break
		}
	}

	// Right iterator
	for ix := x + 1; ix < len(plane[0]); ix++ {
		rightScore = ix - x
		if plane[y][ix] >= height {
			break
		}
	}

	scenicScore = topScore * bottomScore * leftScore * rightScore
	return scenicScore
}
