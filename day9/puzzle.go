package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Simulate your complete hypothetical series of motions. 
How many positions does the tail of the rope visit at 
least once?
`
	positionsVisited := 0
	rope := newBridgeRope()
	for motion := range readMotions() {
		rope.applyMotion(motion)
	}

	solution.Answer = positionsVisited
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = ``
	return solution
}
