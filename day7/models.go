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

		if terminal.curDir != nil {
			terminal.curDir.directories = append(terminal.curDir.directories, directory)
		}
	}

	terminal.curDir = directory
}

// ls will list all files/directories in the current directory. Because our goal is to play back terminal commands that
// have already been run, this really assumes that we have the output of the real terminal ls command. This function
// will take that real output and add to the knowledge of curDir
func (terminal *comTerminal) ls(files []*comFile, dirs []*comDirectory) {
	if terminal.curDir == nil {
		panic("Cannot perform ls, nil directory pointer")
	}
	if len(files) > 0 {
		terminal.curDir.files = append(terminal.curDir.files, files...)
	}
	if len(dirs) > 0 {
		terminal.curDir.directories = append(terminal.curDir.directories, dirs...)
	}
}
