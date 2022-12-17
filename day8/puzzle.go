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
		}
	}

	solution.Answer = amountOfTreesVisibleFromOutside
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
Consider each tree on your map. 
What is the highest scenic score possible 
for any tree?
`
	highestScenicScore := 0

	solution.Answer = highestScenicScore
	return solution
}
