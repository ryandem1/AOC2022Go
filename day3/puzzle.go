package day3

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
)

func Part1() (prioritySum int) {
	priorities := getPriorities()
	priorityToLetter := common.ReverseMap(priorities)

	for rucksack := range readRucksacks() {
		var comp1Occurrences [57]int // We use the priority alphabet map to count occurrences for each item
		var comp2Occurrences [57]int
		var duplicatedLetter rune

		for _, item := range rucksack.compartment1 {
			comp1Occurrences[priorities[item]]++
		}

		for _, item := range rucksack.compartment2 {
			comp2Occurrences[priorities[item]]++
		}

		for i := 0; i < 57; i++ {
			if comp1Occurrences[i] >= 1 && comp2Occurrences[i] >= 1 {
				// Technically, you could just add i to the prioritySum, but this was a little easier to debug
				duplicatedLetter = priorityToLetter[i]
				break
			}
		}
		if duplicatedLetter == 0 {
			panic("Did not find a duplicated letter!")
		}
		prioritySum += priorities[duplicatedLetter]
	}
	return prioritySum
}

// Part2 Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities
// of those item types?
func Part2() (prioritySum int) {
	for rucksacks := range common.ReadChannelInChunks(readRucksacks(), 3) {
		log.Println(rucksacks)
	}
	return prioritySum
}
