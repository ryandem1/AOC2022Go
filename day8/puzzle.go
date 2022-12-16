package day8

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Consider your map; how many trees are visible from 
outside the grid?
`
	amountOfTreesVisibleFromOutside := 0
	trees := buildTreePlane()
	for y, treeRow := range trees {
		for x, height := range treeRow {
			visibleFromOutside := false
			for _, angle := range []viewAngle{top, bottom, left, right} {
				if isVisible(x, y, height, trees, angle) {
					visibleFromOutside = true
					break
				}
			}
			if visibleFromOutside {
				amountOfTreesVisibleFromOutside++
			}
			// fmt.Println(isVisible(x, y, height, trees, top), isVisible(x, y, height, trees, bottom), isVisible(x, y, height, trees, left), isVisible(x, y, height, trees, right))
		}
	}

	solution.Answer = amountOfTreesVisibleFromOutside
	return solution
}
