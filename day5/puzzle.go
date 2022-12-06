package day5

// Part1 After the rearrangement procedure completes, what crate ends up on top of each stack?
func Part1() (topOfEachStack string) {
	stacks := getInitialStacks()
	for action := range readCraneActions() {
		stacks = applyAction(stacks, action)
	}
	for _, stack := range stacks {
		topOfEachStack += string(stack[len(stack)-1])
	}
	return topOfEachStack
}
