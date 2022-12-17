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

// QuadDirection represents linear directions that can be moved on a 2D plane. They are up, down, left, and right.
type QuadDirection int32

const (
	Up QuadDirection = iota
	Down
	Left
	Right
)

// Name will return the string name of a QuadDirection
func (dir *QuadDirection) Name() string {
	name := map[QuadDirection]string{
		Up:    "Up",
		Down:  "Down",
		Left:  "Left",
		Right: "Right",
	}[*dir]

	return name
}

// Coords2D represents a single position on an arbitrary 2D plane
type Coords2D struct {
	X int
	Y int
}

// Move will update a Coords2D X and Y values according to the direction and amount specified.
func (coords *Coords2D) Move(direction QuadDirection, amount int) {
	switch direction {
	case Up:
		coords.Y += amount
	case Down:
		coords.Y -= amount
	case Left:
		coords.X -= 1
	case Right:
		coords.X += 1
	}
}
