package day2

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
	"strings"
)

func winningShape(shape1 handShape, shape2 handShape) *handShape {
	if shape1 == shape2 {
		return nil
	}
	winner := &shape2

	switch shape1 {
	case rock:
		if shape2 == scissors {
			winner = &shape1
		}
	case paper:
		if shape2 == rock {
			winner = &shape1
		}
	case scissors:
		if shape2 == paper {
			winner = &shape1
		}
	default:
		log.Fatal("Unrecognized shape! How is this possible??")
	}
	return winner
}

// getRounds will read data from the input file, format it as a struct and send it through a channel
func getRounds() chan round {
	rounds := make(chan round)
	go func() {
		for line := range common.ReadLinesFromFile("day2") {
			choices := strings.Fields(line)
			r := round{
				opponentShape: letterToHandShape(choices[0]),
				strategyShape: letterToHandShape(choices[1]),
			}
			rounds <- r
		}
		close(rounds)
	}()
	return rounds
}
