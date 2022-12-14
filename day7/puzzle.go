package day7

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Find all the directories with a total size of at most 100000. 
What is the sum of the total sizes of those directories?
`
	totalSize := 0

	terminal := newTerminal()
	terminal.buildDirectoryFromInput()

	for _, dirSize := range getChildDirSizes(terminal.curDir) {
		if dirSize <= 100000 {
			totalSize += dirSize
		}
	}

	solution.Answer = totalSize
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
Find the smallest directory that, if deleted, 
would free up enough space on the filesystem to run the update. 
What is the total size of that directory?
`
	requiredSpace := 30000000
	totalSpace := 70000000

	terminal := newTerminal()
	terminal.buildDirectoryFromInput()
	dirSizes := getChildDirSizes(terminal.curDir)

	unusedSpace := totalSpace - dirSizes["/"]
	spaceNeededToFreeUp := requiredSpace - unusedSpace

	smallestDirSize := dirSizes["/"]
	for _, dirSize := range dirSizes {
		if dirSize >= spaceNeededToFreeUp && dirSize < smallestDirSize {
			smallestDirSize = dirSize
		}
	}

	if smallestDirSize == spaceNeededToFreeUp {
		log.Panicf("Did not find a directory with size > %d", spaceNeededToFreeUp)
	}
	solution.Answer = smallestDirSize
	return solution
}
