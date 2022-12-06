package day5

import "github.com/ryandem1/aoc_2022_go/common"

// Part1 After the rearrangement procedure completes, what crate ends up on top of each stack?
func Part1() (topOfEachStack string) {
	stacks := getInitialStacks()

	for action := range readCraneActions() {
		stacks = applyAction(stacks, action)
	}

	for _, stackLabel := range common.SortedKeys(stacks) {
		stack := stacks[stackLabel]
		topOfEachStack += string(stack[len(stack)-1])
	}

	return topOfEachStack
}
