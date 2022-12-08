package day7

import "github.com/ryandem1/aoc_2022_go/common"

// comFile represents a single file stored on an elf's communication device
type comFile struct {
	name string
	size int
}

// comDirectory represents a single directory on an elf's communication device
type comDirectory struct {
	name        string
	parent      *comDirectory
	directories []*comDirectory
	files       []*comFile
}

// comTerminal represents a single terminal session on an elf's communication device
type comTerminal struct {
	curDir *comDirectory
}

// cd will switch the terminal's context to a different directory. If the directory name is not found in the current
// directory's directories, will add it as a new directory.
func (terminal *comTerminal) cd(dirName string) {
	var directory *comDirectory

	directory = common.FindOneObj(terminal.curDir.directories, func(dir *comDirectory) bool {
		return dir.name == dirName
	})

	if directory == nil {
		directory = &comDirectory{
			name:        dirName,
			parent:      terminal.curDir,
			directories: []*comDirectory{},
			files:       []*comFile{},
		}

		terminal.curDir.directories = append(terminal.curDir.directories, directory)
	}

	terminal.curDir = directory
}
