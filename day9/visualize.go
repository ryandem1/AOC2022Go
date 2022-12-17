package day9

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
)

// VisualizeRopeMovements will visualize how a rope's head and tail positions move across a 2D plane
func VisualizeRopeMovements() {
	var uniqueTailPosVisited []common.Coords2D
	var allPosVisited []common.Coords2D

	rope := newBridgeRope()
	for motion := range readMotions() {
		fmt.Printf("== %s %d ==\n", motion.direction.Name(), motion.amount)

		for i := 0; i < motion.amount; i++ {
			rope.headPos.Move(motion.direction, 1)
			rope.tailPos = moveTailPosition(rope.tailPos, rope.headPos)

			if !common.Contains(uniqueTailPosVisited, rope.tailPos) {
				uniqueTailPosVisited = append(uniqueTailPosVisited, rope.tailPos)
			}

			if !common.Contains(allPosVisited, rope.headPos) {
				allPosVisited = append(allPosVisited, rope.headPos)
			}
			if !common.Contains(allPosVisited, rope.tailPos) {
				allPosVisited = append(allPosVisited, rope.tailPos)
			}
			for _, pos := range allPosVisited {
				if pos == rope.headPos {
					fmt.Println("H")
				} else if pos == rope.tailPos {
					fmt.Println("T")
				} else if common.Contains(uniqueTailPosVisited, pos) {
					fmt.Println("#")
				} else {
					fmt.Println(".")
				}
			}
		}
	}
}
