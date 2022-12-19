package day10

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
	"strconv"
	"strings"
)

// readOperations will read our input file, parse the lines into cpuOperation structs and send them through a channel
func readOperations() chan *cpuOperation {
	operations := make(chan *cpuOperation)

	go func() {
		for line := range common.ReadLinesFromFile("day10") {
			instruction := strings.Fields(line)
			instructionType := instruction[0]

			switch instructionType {
			case "noop":
				operations <- newNoop()
			case "addx":
				amount, err := strconv.Atoi(instruction[1])
				if err != nil {
					panic(err)
				}
				operations <- newAddx(amount)
			default:
				log.Panicf("Unexpected instruction type: %s", instructionType)
			}
		}
		close(operations)
	}()
	return operations
}

// run will execute any cpu operation. Will yield back point-in-time cpuReading objects in an unbuffered channel,
// essentially meaning that each clock cycle is blocking, allowing for analysis one clock cycle at a time
func (c *cpu) run(operations <-chan *cpuOperation) <-chan *cpuReading {
	readings := make(chan *cpuReading)

	go func() {
		for operation := range operations {
			switch operation.opType {
			case noop: // 1 cycle cost
				c.clockCircuit.cycle += 1
				readings <- c.newCPUReading()
			case addx: // 2 cycle cost
				for i := 0; i < 2; i++ {
					c.clockCircuit.cycle += 1
					readings <- c.newCPUReading()
				}
				c.x += operation.V
			}
		}
		close(readings)
	}()
	return readings
}
