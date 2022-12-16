package day8

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
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
