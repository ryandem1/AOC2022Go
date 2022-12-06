package day5

type crateMover uint32

const (
	CrateMover9000 crateMover = iota
	CrateMover9001
)

type CraneAction struct {
	text     string // raw input
	quantity int
	from     string // label of supply stack
	to       string // label of supply stack
}
