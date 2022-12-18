package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"math"
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

// follow will apply the corresponding tail move according to the headPos. While the headPos moves
// according to the motion, the tail position will follow the head according to the puzzle logic.
func follow(tailPos common.Coords2D, headPos common.Coords2D) (movedTailPos common.Coords2D) {
	movedTailPos.X = tailPos.X + int(math.Round(math.Sin(float64(headPos.X-tailPos.X))))
	movedTailPos.Y = tailPos.Y + int(math.Round(math.Sin(float64(headPos.Y-tailPos.Y))))

	if movedTailPos.X == headPos.X && movedTailPos.Y == headPos.Y {
		movedTailPos = tailPos
	}

	return movedTailPos
}

// move will move both the head and tail positions of rope 1 unit in a given direction
func (rope *bridgeRope) move(direction common.QuadDirection) {
	rope.headPos.Move(direction, 1)
	segmentCount := len(rope.ropeSegments)

	for i := 0; i < segmentCount-1; i++ {
		var segmentHead common.Coords2D

		if i == 0 {
			segmentHead = rope.headPos
		} else {
			segmentHead = rope.ropeSegments[i-1]
		}
		rope.ropeSegments[i] = follow(rope.ropeSegments[i], segmentHead)
	}
	segmentHead := rope.headPos
	if segmentCount > 1 {
		segmentHead = rope.ropeSegments[segmentCount-2]
	}

	rope.tailPos = follow(rope.tailPos, segmentHead)
}
