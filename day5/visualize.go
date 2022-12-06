package day5

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
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
	lines[largestStackSize+1] = labelLine
	for _, stackLabel := range common.SortedKeys(stacks) {
		stack := stacks[stackLabel]
		for iStack := largestStackSize; iStack >= 0; iStack-- {
			if iStack > len(stack)-1 {
				lines[iStack] += "   "
			} else {
				lines[iStack] += "[" + string(stack[iStack]) + "]"
			}
			lines[iStack] += " "
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
