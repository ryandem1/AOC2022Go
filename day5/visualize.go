package day5

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"strings"
)

// prettyPrintStacks will print out the stacks that are used in processing in the original format from the prompt
func prettyPrintStacks(stacks map[string]string) {
	largestStackSize := 0
	labelLine := ""

	for _, stackLabel := range common.SortedKeys(stacks) {
		stackSize := len(stacks[stackLabel])
		if stackSize > largestStackSize {
			largestStackSize = stackSize
		}
		labelLine += " " + string(stackLabel) + "  "
	}

	lines := make([]string, largestStackSize+2)
	for _, stackLabel := range common.SortedKeys(stacks) {
		stack := stacks[stackLabel]
		for iStack := largestStackSize + 1; iStack >= 0; iStack-- {
			if iStack > len(stack)-1 {
				lines[iStack] += "   "
			} else {
				lines[iStack] += "[" + string(stack[iStack]) + "]"
			}
			lines[iStack] += " "
		}
	}
	lines = append([]string{labelLine}, lines...) // Prepend label line
	for iLines := len(lines) - 1; iLines >= 0; iLines-- {
		fmt.Println(lines[iLines])
	}
}

// Visualize will provide an interactive supplyStacks supply that can visualize how the stack changes with each new
// crane action. Must select which mover spec (either the 9000 or 9001), because it changes how actions are applied
func Visualize(mover crateMover) {
	stacks := getInitialStacks()
	fmt.Println("INITIAL STACK:")
	prettyPrintStacks(stacks)

	for action := range readCraneActions() {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Println("ACTION: " + action.text)

		stacks = applyAction(stacks, action, mover)
		prettyPrintStacks(stacks)
	}
}
