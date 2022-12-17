package day9

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"sort"
)

// VisualizeRopeMovements will visualize how a rope's head and tail positions move across a 2D plane
func VisualizeRopeMovements() {
	var uniqueTailPosVisited []common.Coords2D
	var allPosVisited []common.Coords2D

	rope := newBridgeRope()
	fmt.Println("== Begin ==")
	fmt.Println("H") // Begins with head overlapping tail at position S

	for motion := range readMotions() {
		fmt.Printf("== %s %d ==\n", motion.direction.Name(), motion.amount)

		for i := 0; i < motion.amount; i++ {
			plane := "" // output to be printed

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

			// Sorts all positions by y/x asc
			sort.Slice(allPosVisited, func(i int, j int) bool {
				return allPosVisited[i].Y < allPosVisited[j].Y || allPosVisited[i].X < allPosVisited[j].X
			})

			for y := allPosVisited[len(allPosVisited)-1].Y + 10; y >= allPosVisited[0].Y-10; y-- {
				for x := allPosVisited[0].X - 20; x <= allPosVisited[len(allPosVisited)-1].X+20; x++ {
					pos := common.Coords2D{
						X: x,
						Y: y,
					}

					if x == rope.headPos.X && y == rope.headPos.Y {
						plane += "H"
					} else if x == rope.tailPos.X && y == rope.tailPos.Y {
						plane += "T"
					} else if common.Contains(uniqueTailPosVisited, pos) {
						plane += "#"
					} else {
						plane += "."
					}
				}
				plane += "\n"
			}
			fmt.Println(plane)
		}
	}
}
