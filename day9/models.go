package day9

import "github.com/ryandem1/aoc_2022_go/common"

// motionDirection are the various directions that the head of the bridgeRope can move
type motionDirection int32

const (
	up motionDirection = iota
	down
	left
	right
)

// bridgeRope represents a single rope on a 2d plane with a head knot and a tail knot
type bridgeRope struct {
	headPos        common.Coords2D
	tailPos        common.Coords2D
	tailPosVisited []common.Coords2D
}

// newBridgeRope will initialize a new bridgeRope at starting position S, which will be set as 0, 0. With this config
// there will be negative coordinates
func newBridgeRope() *bridgeRope {
	rope := &bridgeRope{
		headPos: common.Coords2D{
			X: 0,
			Y: 0,
		},
		tailPos: common.Coords2D{
			X: 0,
			Y: 0,
		},
		tailPosVisited: []common.Coords2D{},
	}
	return rope
}

// ropeMotion is our input, it is a single movement of a fixed length in a specific motionDirection
type ropeMotion struct {
	direction motionDirection
	amount    int
}
