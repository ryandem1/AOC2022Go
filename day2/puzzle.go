package day2

import "github.com/ryandem1/aoc_2022_go/common"

// Part1 What would your total score be if everything goes exactly according to your strategy guide?
func Part1() (totalScore int) {
	for round := range getRounds(common.Part1) {
		totalScore += scoreRound(round)
	}
	return totalScore
}

// Part2 What would your total score be if everything goes exactly according to your strategy guide?
func Part2() (totalScore int) {
	for round := range getRounds(common.Part2) {
		totalScore += scoreRound(round)
	}
	return totalScore
}
