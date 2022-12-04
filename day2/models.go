package day2

import (
	"fmt"
)

type handShape int32

const (
	rock handShape = iota
	paper
	scissors
)

type day2Part int32

const (
	part1 day2Part = iota
	part2
)

// letterToHandShapePart1 will convert the letter from the input file to a handShape according to Part1 logic
func letterToHandShapePart1(letter string) handShape {
	switch letter {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		panic(fmt.Sprintf("Unhandled letter %s", letter))
	}
}

// letterToHandShapePart2 will convert the letter from the input file to a handShape according to Part2 logic. What
// this will really do is basically translates the part2 letters to the part1 letters so the rest of the functions work
func letterToHandShapePart2(letter string, opponentHandShape handShape) (outShape handShape) {
	switch letter {
	case "A", "B", "C": // Only need to translate XYZ... ABC definition has not changed
		outShape = letterToHandShapePart1(letter)
	case "X": // Return the losing shape
		for _, shape := range []handShape{rock, paper, scissors} {
			if winningShape(&opponentHandShape, &shape) == &opponentHandShape {
				outShape = shape
			}
		}
	case "Y": // Tie
		outShape = opponentHandShape
	case "Z": // Return winning shape
		for _, shape := range []handShape{rock, paper, scissors} {
			if winningShape(&opponentHandShape, &shape) == &shape {
				outShape = shape
			}
		}
	default:
		panic(fmt.Sprintf("Unhandled letter %s", letter))
	}
	return outShape
}

// RPSRound represents a single round of rock-paper-scissors between the opponent and strategy
type RPSRound struct {
	opponentShape handShape
	strategyShape handShape
}
