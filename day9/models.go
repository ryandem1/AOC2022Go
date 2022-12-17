package day9

type motionDirection int32

const (
	up motionDirection = iota
	down
	left
	right
)

type coords struct {
	x int
	y int
}

type bridgeRope struct {
	headPos coords
	tailPos coords
}

type ropeMotion struct {
	direction motionDirection
	amount    int
}
