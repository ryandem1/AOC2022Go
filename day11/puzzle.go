package day11

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Figure out which monkeys to chase by counting how many 
items they inspect over 20 rounds. 

What is the level of monkey business after 20 rounds 
of stuff-slinging simian shenanigans?
`

	monkeys := readMonkeys()
	monkey1 := monkeys[0]
	// monkey2 := monkeys[1]

	fmt.Println(monkey1.test(monkey1.items[0]))
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = ``
	return solution
}
