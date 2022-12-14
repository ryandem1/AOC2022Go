package day2

import "github.com/ryandem1/aoc_2022_go/common"

func Part1() (solution common.Solution) {
	solution.Prompt = `
What would your total score be if everything goes 
exactly according to your strategy guide?
`
	var totalScore int

	for round := range getRounds(common.Part1) {
		totalScore += scoreRound(round)
	}
	solution.Answer = totalScore
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
What would your total score be if everything goes 
exactly according to your strategy guide?
`
	var totalScore int

	for round := range getRounds(common.Part2) {
		totalScore += scoreRound(round)
	}
	solution.Answer = totalScore
	return solution
}
