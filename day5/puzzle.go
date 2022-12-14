package day5

import "github.com/ryandem1/aoc_2022_go/common"

func Part1() (solution common.Solution) {
	solution.Prompt = `
After the rearrangement procedure completes, 
what crate ends up on top of each stack?
`
	var topOfEachStack string

	stacks := getInitialStacks()

	for action := range readCraneActions() {
		stacks = applyAction(stacks, action, CrateMover9000)
	}

	for _, stackLabel := range common.SortedKeys(stacks) {
		stack := stacks[stackLabel]
		topOfEachStack += string(stack[len(stack)-1])
	}

	solution.Answer = topOfEachStack
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
After the rearrangement procedure completes, 
what crate ends up on top of each stack?
`
	var topOfEachStack string
	stacks := getInitialStacks()

	for action := range readCraneActions() {
		stacks = applyAction(stacks, action, CrateMover9001)
	}

	for _, stackLabel := range common.SortedKeys(stacks) {
		stack := stacks[stackLabel]
		topOfEachStack += string(stack[len(stack)-1])
	}

	solution.Answer = topOfEachStack
	return solution
}
