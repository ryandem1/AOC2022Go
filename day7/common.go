package day7

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
	"strings"
)

// addNewDirIfNotExists will add a new directory (initialized by dirName) to the terminal's current directory if it is
// not already there (checks by name). Will also return the directory either retrieved or created.
func (terminal *comTerminal) addNewDirIfNotExists(dirName string) (directory *comDirectory) {
	if terminal.curDir != nil {
		directory = common.FindOneObj(terminal.curDir.directories, func(dir *comDirectory) bool {
			return dir.name == dirName
		})
	}

	if directory == nil {
		directory = &comDirectory{
			name:        dirName,
			parent:      terminal.curDir,
			directories: []*comDirectory{},
			files:       []*comFile{},
		}

		if terminal.curDir != nil {
			terminal.curDir.directories = append(terminal.curDir.directories, directory)
		}
	}
	return directory
}

// cd will switch the terminal's context to a different directory. If the directory name is not found in the current
// directory's directories, will add it as a new directory.
func (terminal *comTerminal) cd(dirName string) {
	var directory *comDirectory

	if dirName == ".." {
		if terminal.curDir.parent == nil {
			panic("Cannot go up a directory, already at the root!")
		}
		terminal.curDir = terminal.curDir.parent
		return
	}

	directory = terminal.addNewDirIfNotExists(dirName)
	terminal.curDir = directory
}

// ls will list all files/directories in the current directory. Because our goal is to play back terminal commands that
// have already been run, this really assumes that we have the output of the real terminal ls command. This function
// will take that real output and add to the knowledge of curDir
func (terminal *comTerminal) ls(files []*comFile, dirNames []string) {
	if terminal.curDir == nil {
		panic("Cannot perform ls, nil directory pointer")
	}
	if len(files) > 0 {
		terminal.curDir.files = append(terminal.curDir.files, files...)
	}
	for _, dir := range dirNames {
		terminal.addNewDirIfNotExists(dir)
	}
}

// executeCommand will read from a string channel of the command lines. Will return ok if a command was executed, or
// else it will return false, indicating that there are no more commands to read from the input.
func (terminal *comTerminal) executeCommand(cmdOutput chan string) *ExeReceipt {
	// Need to add a buffer because we will be placing a single value back into the channel
	cmdLine, ok := <-cmdOutput
	receipt := &ExeReceipt{
		nextCommand: "",
		ok:          ok,
	}
	if !ok {
		return receipt
	}

	cmd := strings.Fields(cmdLine)

	if cmd[0] != "$" {
		panic("Next line is not a command! Something is wrong..." + fmt.Sprintf("Line: %s", cmdLine))
	}

	commandType := cmd[1]

	switch commandType {
	case "cd":
		targetDir := cmd[2]
		terminal.cd(targetDir)
	case "ls":
		var dirNames []string
		var files []*comFile

		for lsLine := range cmdOutput {
			lsItem := strings.Fields(lsLine)

			if lsItem[0] == "$" {
				// We have reach the next command, so our ls output is done, we can include the next command in the
				// receipt
				receipt.nextCommand = lsLine
				break
			} else if lsItem[0] == "dir" {
				dirName := lsItem[1]
				dirNames = append(dirNames, dirName)
			} else { // Must be a file
				fileSize, err := strconv.Atoi(lsItem[0])
				if err != nil {
					panic(err)
				}
				fileName := lsItem[1]
				files = append(files, &comFile{
					name: fileName,
					size: fileSize,
				})
			}
		}
		terminal.ls(files, dirNames)
	default:
		panic(fmt.Sprintf("Unhandled/invalid command type: %s", commandType))
	}
	return receipt
}

// buildDirectoryFromInput will execute all commands from the day7 input file and will build the directory structure
// the terminal's curDir will be reset to the root "/"
func (terminal *comTerminal) buildDirectoryFromInput() {
	lineReader := common.ReadLinesFromFile("day7")
	cmdBuf := make(chan string) // To store the ls commands

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
}

// getChildDirSizes will return a map of directories (by name) and their sizes (int) contained in a directory
func getChildDirSizes(dir *comDirectory) map[string]int {
	dirSizes := make(map[string]int)
	dirSizes[dir.path()] = dir.totalSize()

	for _, childDir := range dir.directories {
		childDirSizes := getChildDirSizes(childDir)

		for dirName, dirSize := range childDirSizes {
			dirSizes[dirName] = dirSize
		}
	}
	return dirSizes
}
