package day7

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
