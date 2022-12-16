package day8

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
)

// buildTreeMap will read the input and build the 2D map of trees
func buildTreeMap() map[string]*tree {
	treeMap := make(map[string]*tree)

	y := 0
	for line := range common.ReadLinesFromFile("day8") {
		for x, rawHeight := range line {
			height, err := strconv.Atoi(string(rawHeight))
			if err != nil {
				panic(err)
			}

			t := &tree{
				x:      x,
				y:      y,
				height: height,
			}
			treeMap[t.coords()] = t
		}
		y++
	}
	return treeMap
}
