package day5

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strings"
)

// getInitialStacks will read the header of the input file and parse it to get the initial state of all stacks
func getInitialStacks() map[string]string {
	stacks := make(map[string]string)

	// Front load the header because we really want to read from the bottom of the header up
	var headerLines []string // header lines
	for line := range common.ReadLinesFromFile("day5") {
		if line == "" {
			break
		}
		headerLines = append(headerLines, line)
	}

	stackLabels := strings.Fields(headerLines[len(headerLines)-1])
	for _, label := range stackLabels {
		stacks[label] = ""
	}

	// Read from bottom to top of stack
	for iHeaderLine := len(headerLines) - 2; iHeaderLine >= 0; iHeaderLine-- {
		line := headerLines[iHeaderLine]

		iStack := 0

		for iLine := 1; iLine < len(line); iLine += 4 {
			stacks[stackLabels[iStack]] += string(line[iLine])
			iStack++
		}
	}

	return stacks
}
