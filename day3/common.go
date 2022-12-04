package day3

import "github.com/ryandem1/aoc_2022_go/common"

// readRucksacks will parse the input file for elfRucksack objects and send them through a channel
func readRucksacks() chan elfRucksack {
	rucksacks := make(chan elfRucksack)

	go func() {
		for line := range common.ReadLinesFromFile("day3") {
			rucksack := newRucksack(line)
			rucksacks <- rucksack
		}
		close(rucksacks)
	}()
	return rucksacks
}
