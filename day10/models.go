package day10

import (
	"github.com/ryandem1/aoc_2022_go/common"
)

type cpuOperationType int32

const (
	noop cpuOperationType = iota
	addx
)

type cpuOperation struct {
	opType cpuOperationType
	V      int // Only applicable to addx instructions
}

// newNoop takes one cycle to complete. It has no other effect.
func newNoop() *cpuOperation {
	op := &cpuOperation{
		opType: noop,
		V:      0,
	}
	return op
}

// newAddx addx V takes two cycles to complete. After two cycles, the X register is increased by the value V.
// (V can be negative.)
func newAddx(V int) *cpuOperation {
	op := &cpuOperation{
		opType: addx,
		V:      V,
	}
	return op
}

// cpuClockCircuit represents the clock unit that increments the cpu cycle at a constant rate.
type cpuClockCircuit struct {
	cycle int
}

// cpu represents the simple CPU as laid out in the puzzle. Has one x register that contains a signed int value. Also
// has a clockCircuit unit that can keep track of the cpu cycles
type cpu struct {
	x            int
	clockCircuit *cpuClockCircuit
}

// signalStrength is the cycle number of the cpuClockCircuit multiplied by the value of the x register
func (c *cpu) signalStrength() int {
	return c.x * c.clockCircuit.cycle
}

// newCPU initializes a new cpu with a cpuCLockCircuit
func newCPU() *cpu {
	cc := &cpuClockCircuit{cycle: 0}
	c := &cpu{
		x:            1,
		clockCircuit: cc,
	}
	return c
}

// cpuReading is a point-in-time reading of a cpu's signalStrength and the current cycle of its cpuClockCircuit
type cpuReading struct {
	signalStrength int
	xRegisterValue int
	cycle          int
}

// newCPUReading will take a point-in-time reading of a CPU
func (c *cpu) newCPUReading() *cpuReading {
	reading := &cpuReading{
		signalStrength: c.signalStrength(),
		xRegisterValue: c.x,
		cycle:          c.clockCircuit.cycle,
	}
	return reading
}

// crtDisplay represents a single display of a communication device
type crtDisplay struct {
	pixels       []common.Coords2D
	clockCircuit *cpuClockCircuit
	sprite       [3]int // X positions of sprite
}

// newCRTDisplay will create a new crtDisplay. There must be an equal number of pixels per row. Takes in a clock so
// that it can be synchronized with a cpu
func newCRTDisplay(height int, width int, clock *cpuClockCircuit) *crtDisplay {
	var pixels []common.Coords2D

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixels = append(pixels, common.Coords2D{X: x, Y: y})
		}
	}
	display := &crtDisplay{
		pixels:       pixels,
		clockCircuit: clock,
		sprite:       [3]int{0, 1, 2},
	}
	return display
}

// moveSprite will move a crtDisplay's sprite to a certain x position. The x position will correlate to the middle
// of the sprite
func (display *crtDisplay) moveSprite(x int) {
	display.sprite = [3]int{x - 1, x, x + 1}
}
