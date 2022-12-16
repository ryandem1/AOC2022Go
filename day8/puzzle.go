package day8

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Consider your map; how many trees are visible from 
outside the grid?
`

	trees := buildTreePlane()
	for y, treeRow := range trees {
		for x, height := range treeRow {
			fmt.Println(isVisible(x, y, height, trees, top), isVisible(x, y, height, trees, bottom), isVisible(x, y, height, trees, left), isVisible(x, y, height, trees, right))
		}
	}
	return solution
}
