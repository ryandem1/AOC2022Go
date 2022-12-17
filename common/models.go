package common

// DayPart represents the different parts to a puzzle
type DayPart int32

const (
	Part1 DayPart = iota
	Part2
)

// Solution objects represent a single solution to a DayPart
type Solution struct {
	Prompt string
	Answer any
}

// Coords2D represents a single position on an arbitrary 2D plane
type Coords2D struct {
	X int
	Y int
}
