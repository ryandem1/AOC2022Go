package common

type DayPart int32

const (
	Part1 DayPart = iota
	Part2
)

type Solution struct {
	Prompt string
	Answer any
}
