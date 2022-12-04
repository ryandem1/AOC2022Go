package day2

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
	"strings"
)

func winningShape(shape1 *handShape, shape2 *handShape) *handShape {
	if *shape1 == *shape2 {
		return nil
	}
	winner := shape2

	switch *shape1 {
	case rock:
		if *shape2 == scissors {
			winner = shape1
		}
	case paper:
		if *shape2 == rock {
			winner = shape1
		}
	case scissors:
		if *shape2 == paper {
			winner = shape1
		}
	default:
		log.Fatal("Unrecognized shape! How is this possible??")
	}
	return winner
}

// getRounds will read data from the input file, format it as a struct and send it through a channel
func getRounds(part common.DayPart) chan RPSRound {
	rounds := make(chan RPSRound)
	go func() {
		for line := range common.ReadLinesFromFile("day2") {
			choices := strings.Fields(line)

			round := RPSRound{}
			switch part {
			case common.Part1:
				round = RPSRound{
					opponentShape: letterToHandShapePart1(choices[0]),
					strategyShape: letterToHandShapePart1(choices[1]),
				}
			case common.Part2:
				opponentShape := letterToHandShapePart1(choices[0])
				round = RPSRound{
					opponentShape: opponentShape,
					strategyShape: letterToHandShapePart2(choices[1], opponentShape),
				}
			default:
				panic(fmt.Sprintf("Unrecognized part: %d", part))
			}
			rounds <- round
		}
		close(rounds)
	}()
	return rounds
}

// scoreRound will rake in a round and perform the scoring logic to output an int score. Scoring logic is as follows:
//
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func scoreRound(round RPSRound) (score int) {
	switch round.strategyShape {
	case rock:
		score += 1
	case paper:
		score += 2
	case scissors:
		score += 3
	}

	winner := winningShape(&round.strategyShape, &round.opponentShape)
	switch winner {
	case nil: // tie
		score += 3
	case &round.strategyShape:
		score += 6
	case &round.opponentShape:
		score += 0
	}

	return score
}
