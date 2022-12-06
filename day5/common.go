package day5

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"sort"
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
	sort.Strings(stackLabels)

	for _, label := range stackLabels {
		stacks[label] = ""
	}

	// Read from bottom to top of stack
	for iHeaderLine := len(headerLines) - 2; iHeaderLine >= 0; iHeaderLine-- {
		line := headerLines[iHeaderLine]

		iStack := 0

		for iLine := 1; iLine < len(line); iLine += 4 {
			value := string(line[iLine])
			if value != " " {
				stacks[stackLabels[iStack]] += value
			}
			iStack++
		}
	}

	return stacks
}

// readCraneActions will read the input file from the header down, parse the actions into the struct and send it
// through the channel
func readCraneActions() chan CraneAction {
	actions := make(chan CraneAction)

	go func() {
		pastHeader := false
		for line := range common.ReadLinesFromFile("day5") {
			if line == "" { // Blank line determines end of header
				pastHeader = true
				continue
			}
			if !pastHeader {
				continue
			}
			var quantity int
			var to string
			var from string

			numParsed, err := fmt.Sscanf(line, "move %d from %s to %s", &quantity, &from, &to)
			if err != nil {
				panic(err)
			}
			if numParsed != 3 {
				panic(fmt.Sprintf("Could not parse line! Line: %s", line))
			}
			action := CraneAction{
				quantity: quantity,
				from:     from,
				to:       to,
			}
			actions <- action
		}
		close(actions)
	}()
	return actions
}

// applyAction will apply an action on a map of supplyStacks. Will return the altered stack
func applyAction(supplyStacks map[string]string, action CraneAction) map[string]string {
	fStack := supplyStacks[action.from]
	tStack := supplyStacks[action.to]

	supplies := fStack[len(fStack)-action.quantity:] // Take off the top of from stack
	fStack = fStack[:len(fStack)-action.quantity]

	tStack += supplies // Add to the destination stack
	supplyStacks[action.from] = fStack
	supplyStacks[action.to] = tStack
	return supplyStacks
}
