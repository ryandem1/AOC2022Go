package day6

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

// Part1 How many characters need to be processed before the first start-of-packet marker is detected?
func Part1() (numCharacters int) {
	line := <-common.ReadLinesFromFile("day6")
	numCharacters += 4
	found := false

	for window := range common.ReadSliceWithWindow([]rune(line), 4) {
		if len(window) < 4 {
			continue // window must have 4 values to be considered
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
