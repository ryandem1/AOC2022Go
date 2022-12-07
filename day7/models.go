package day7

// comTerminalCommand is an enum type that represent the different types of commands you can perform on an elf's
// communication device's terminal
type comTerminalCommand int32

const (
	cd comTerminalCommand = iota
	ls
)

// comFile represents a single file stored on an elf's communication device
type comFile struct {
	name string
	size int
}

// comDirectory represents a single directory on an elf's communication device
type comDirectory struct {
	directories []*comDirectory
	files       []*comFile
}

// commandExecution represents a single execution of a command, part of the puzzle input
type commandExecution struct {
	command  comTerminalCommand
	argument string
}
