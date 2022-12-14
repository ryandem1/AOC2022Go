package day4

import "github.com/ryandem1/aoc_2022_go/common"

func Part1() (solution common.Solution) {
	solution.Prompt = `
In how many assignment pairs does one range 
fully contain the other?
`
	var fullyOverlappingPairsCount int

	for pair := range readCleaningAssignmentPairs() {
		for i := 0; i < 2; i++ {
			if pair[i].start <= pair[i^1].start && pair[i].end >= pair[i^1].end {
				fullyOverlappingPairsCount++
				break
			}
		}
	}
	solution.Answer = fullyOverlappingPairsCount
	return solution
}

// Part2 In how many assignment pairs do the ranges overlap?
func Part2() (solution common.Solution) {
	solution.Prompt = `
In how many assignment pairs do the ranges 
overlap?
`
	var overlappingPairsCount int

	for pair := range readCleaningAssignmentPairs() {
		for i := 0; i < 2; i++ {
			startInPairAssignment := pair[i^1].end >= pair[i].start && pair[i].start >= pair[i^1].start
			endInPairAssignment := pair[i^1].end >= pair[i].end && pair[i].end >= pair[i^1].start

			if startInPairAssignment || endInPairAssignment {
				overlappingPairsCount++
				break
			}
		}
	}
	solution.Answer = overlappingPairsCount
	return solution
}
