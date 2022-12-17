package day9

import "github.com/ryandem1/aoc_2022_go/common"

// bridgeRope represents a single rope on a 2d plane with a head knot and a tail knot. Between the head and tail knot
// can be an arbitrary number of rope segments
type bridgeRope struct {
	headPos      common.Coords2D
	ropeSegments []common.Coords2D
	tailPos      common.Coords2D
}

// newBridgeRope will initialize a new bridgeRope at starting position S, which will be set as 0, 0. With this config
// there will be negative coordinates
// segCount defines how many rope segments are between the head knot and the tail knot
func newBridgeRope(segCount int) *bridgeRope {
	sPos := common.Coords2D{
		X: 0,
		Y: 0,
	}
	var segments []common.Coords2D
	for i := 0; i < segCount; i++ {
		segments = append(segments, sPos)
	}

	rope := &bridgeRope{
		headPos:      sPos,
		ropeSegments: segments,
		tailPos:      sPos,
	}
	return rope
}

// ropeMotion is our input, it is a single movement of a fixed length in a specific motionDirection
type ropeMotion struct {
	direction common.QuadDirection
	amount    int
}
