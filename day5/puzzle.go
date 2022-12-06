package day5

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
)

// Part1 After the rearrangement procedure completes, what crate ends up on top of each stack?
func Part1() (topOfEachStack string) {
	common.PrettyPrintMap(getInitialStacks())
	for action := range readCraneActions() {
		fmt.Println(action)
	}
	return topOfEachStack
}
