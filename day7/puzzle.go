package day7

import "github.com/ryandem1/aoc_2022_go/common"

func Part1() (solution common.Solution) {
	solution.Prompt = `
Find all the directories with a total size of at most 100000. 
What is the sum of the total sizes of those directories?
`
	totalSize := 0

	terminal := newTerminal()
	dirSizes := make(map[string]int) // comDir name, comDir total size

	terminal.buildDirectoryFromInput()
	dirSizes = getChildDirSizes(terminal.curDir)

	for _, dirSize := range dirSizes {
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
	solution.Answer = 0
	return solution
}
