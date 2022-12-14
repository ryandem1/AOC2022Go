package day6

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

func Day6(part common.DayPart) (solution common.Solution) {
	if part == common.Part1 {
		solution.Prompt = `
How many characters need to be processed before the 
first start-of-packet marker is detected?
`
	} else {
		solution.Prompt = `
How many characters need to be processed before the 
first start-of-packet marker is detected?
`
	}
	numCharacters := 0
	distinctCharsNeededForSignal := 4
	if part == common.Part2 {
		distinctCharsNeededForSignal = 14
	}

	line := <-common.ReadLinesFromFile("day6")
	numCharacters += distinctCharsNeededForSignal
	found := false

	for window := range common.ReadSliceWithWindow([]rune(line), distinctCharsNeededForSignal) {
		if len(window) < distinctCharsNeededForSignal {
			continue // window must have 4 or 14 values to be considered
		}

		charOccurrences := make(map[rune]int)
		for _, char := range window {
			if _, ok := charOccurrences[char]; ok {
				charOccurrences[char]++
			} else {
				charOccurrences[char] = 1
			}
		}

		for _, occurrences := range charOccurrences {
			if occurrences == 1 {
				found = true
			} else {
				found = false
				break
			}
		}

		if found {
			solution.Answer = numCharacters
			return solution
		}
		numCharacters++
	}
	panic("Did not find the start-of-packet marker!")
}
