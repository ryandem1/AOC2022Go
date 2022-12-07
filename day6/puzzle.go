package day6

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

// Day6
// Part 1: How many characters need to be processed before the first start-of-packet marker is detected?
// Part 2: How many characters need to be processed before the first start-of-message marker is detected?
func Day6(part common.DayPart) (numCharacters int) {
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
			return numCharacters
		}
		numCharacters++
	}
	panic("Did not find the start-of-packet marker!")
}
