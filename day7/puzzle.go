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
	dirSizes := make(map[string]int) // comDir name, comDir total size
	cmdBuf := make(chan string)      // To store the ls commands

	receipt := &ExeReceipt{
		nextCommand: "",
		ok:          true,
	}
	for receipt.ok {
		if receipt.nextCommand != "" {
			go func() {
				cmdBuf <- receipt.nextCommand
			}()
			receipt = terminal.executeCommand(cmdBuf)
		} else {
			receipt = terminal.executeCommand(lineReader)
		}
	}
	close(cmdBuf)

	for terminal.curDir.name != "/" {
		terminal.cd("..")
	}
	dirSizes = getChildDirSizes(terminal.curDir)

	for dirName, dirSize := range dirSizes {
		fmt.Println(dirName, dirSize)
		if dirSize <= 100000 {
			totalSize += dirSize
		}
	}
	return totalSize
}
