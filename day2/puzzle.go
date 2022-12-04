package day2

import "log"

// Part1 What would your total score be if everything goes exactly according to your strategy guide?
func Part1() (totalScore int) {
	for round := range getRounds() {
		log.Println(round)
	}
	return totalScore
}
