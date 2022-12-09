package day7

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
)

// Part1 Find all the directories with a total size of at most 100000. What is the sum of the total sizes of those
// directories?
func Part1() (totalSize int) {
	terminal := newTerminal()
	lineReader := common.ReadLinesFromFile("day7")

	ok := true
	for ok {
		ok = terminal.executeCommand(lineReader)
		fmt.Println(terminal.curDir)
	}
	close(lineReader)
	return totalSize
}
