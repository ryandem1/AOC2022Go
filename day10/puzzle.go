package day10

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Find the signal strength during the 20th, 60th, 
100th, 140th, 180th, and 220th cycles. 

What is the sum of these six signal strengths?
`
	signalStrengthSum := 0
	c := newCPU()

	for reading := range c.run(readOperations()) {
		if common.Contains([]int{20, 60, 100, 140, 180, 220}, reading.cycle) {
			signalStrengthSum += reading.signalStrength
		}
	}

	solution.Answer = signalStrengthSum
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `

`
	signalStrengthSum := 0

	solution.Answer = signalStrengthSum
	return solution
}
