package day7

// Part1 Find all the directories with a total size of at most 100000. What is the sum of the total sizes of those
// directories?
func Part1() (totalSize int) {
	terminal := newTerminal()
	dirSizes := make(map[string]int) // comDir name, comDir total size

	terminal.buildDirectoryFromInput()
	dirSizes = getChildDirSizes(terminal.curDir)

	for _, dirSize := range dirSizes {
		if dirSize <= 100000 {
			totalSize += dirSize
		}
	}
	return totalSize
}

func Part2() (totalSize int) {
	return totalSize
}
