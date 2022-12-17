package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
	"strings"
)

// readMotions will convert the raw string line input into ropeMotion objects and send them through a channel for use
// in the puzzle
func readMotions() chan *ropeMotion {
	motions := make(chan *ropeMotion)

	go func() {
		for line := range common.ReadLinesFromFile("day9") {
			fields := strings.Fields(line)

			direction := map[string]common.QuadDirection{
				"U": common.Up,
				"D": common.Down,
				"L": common.Left,
				"R": common.Right,
			}[fields[0]]
			amount, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}

			motion := &ropeMotion{
				direction: direction,
				amount:    amount,
			}
			motions <- motion
		}
		close(motions)
	}()
	return motions
}

// moveTailPosition will apply the corresponding tail move according to the headPos. While the headPos moves
// according to the motion, the tail position will follow the head according to the puzzle logic.
func moveTailPosition(tailPos common.Coords2D, headPos common.Coords2D) (movedTailPos common.Coords2D) {
	xDelta := headPos.X - tailPos.X
	yDelta := headPos.Y - tailPos.Y
	if xDelta > 1 {
		movedTailPos.X = headPos.X - tailPos.X - 1
	}
	if yDelta > 1 {
		movedTailPos.Y = headPos.Y - tailPos.Y - 1
	}
	return movedTailPos
}

// move will move both the head and tail positions of rope 1 unit in a given direction
func (rope *bridgeRope) move(direction common.QuadDirection) {
	rope.headPos.Move(direction, 1)
	rope.tailPos = moveTailPosition(rope.tailPos, rope.headPos)
}
