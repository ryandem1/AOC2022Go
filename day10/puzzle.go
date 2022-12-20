package day10

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"sort"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Find the signal strength during the 20th, 60th, 
100th, 140th, 180th, and 220th cycles. 

What is the sum of these six signal strengths?
`
	signalStrengthSum := 0
	c := newCPU()

	for reading := range c.run(readOperations()) {
		if common.Contains([]int{20, 60, 100, 140, 180, 220}, reading.cycle) {
			signalStrengthSum += reading.signalStrength
		}
	}

	solution.Answer = signalStrengthSum
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
Render the image given by your program. 

What eight capital letters appear on your CRT?
`
	letters := ""
	c := newCPU()
	displayHeight := 6
	displayWidth := 40
	display := newCRTDisplay(displayHeight, displayWidth, c.clockCircuit)
	sort.Slice(display.pixels, func(i int, j int) bool {
		return display.pixels[i].X < display.pixels[j].X && display.pixels[i].Y < display.pixels[j].Y
	})

	iPixel := 0
	letters += "\n"
	for reading := range c.run(readOperations()) {
		display.moveSprite(reading.xRegisterValue)
		if iPixel > 0 && iPixel%displayWidth == 0 {
			letters += "\n"
		}
		if common.Contains(display.sprite[:], display.pixels[iPixel].X) {
			letters += "#"
		} else {
			letters += "."
		}
		iPixel++
	}
	letters += "\n"

	solution.Answer = letters
	return solution
}
