package day9

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"sort"
	"strconv"
)

// VisualizeRopeMovements will visualize how a rope's head and tail positions move across a 2D plane
func VisualizeRopeMovements(ropeSize int) {
	var uniqueTailPosVisited []common.Coords2D
	var allPosVisited []common.Coords2D

	rope := newBridgeRope(ropeSize - 2)
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

			// Sorts all positions by y to get Y bounds
			sort.Slice(allPosVisited, func(i int, j int) bool {
				return allPosVisited[i].Y < allPosVisited[j].Y
			})
			lowestY := allPosVisited[0].Y
			highestY := allPosVisited[len(allPosVisited)-1].Y

			// Sorts by X for x bounds
			sort.Slice(allPosVisited, func(i int, j int) bool {
				return allPosVisited[i].X < allPosVisited[j].X
			})
			lowestX := allPosVisited[0].X
			highestX := allPosVisited[len(allPosVisited)-1].X

			for y := highestY; y >= lowestY; y-- {
				for x := lowestX; x <= highestX; x++ {
					pos := common.Coords2D{
						X: x,
						Y: y,
					}

					segmentIndex := common.FindIndex(rope.ropeSegments, pos)
					if x == rope.headPos.X && y == rope.headPos.Y {
						plane += "H"
					} else if x == rope.tailPos.X && y == rope.tailPos.Y {
						plane += "T"
					} else if segmentIndex > -1 { // If there are segments, we prefer to write those first
						plane += strconv.Itoa(segmentIndex)
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
