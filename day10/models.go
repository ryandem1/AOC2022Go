package day10

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

// newCPU initializes a new cpu with a cpuCLockCircuit
func newCPU() *cpu {
	cc := &cpuClockCircuit{cycle: 0}
	c := &cpu{
		x:            0,
		clockCircuit: cc,
	}
	return c
}
