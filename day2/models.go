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

// letterToHandShape will convert the letter from the input file to a handShape
func letterToHandShape(letter string) handShape {
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

// RPSRound represents a single round of rock-paper-scissors between the opponent and strategy
type RPSRound struct {
	opponentShape handShape
	strategyShape handShape
}
