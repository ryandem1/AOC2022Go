package day1

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
)

// readElves will read lines from the input file, deserialize the data into an Elf and yield it
func readElves() chan Elf {
	elves := make(chan Elf)
	go func() {
		var foods []Food

		for line := range common.ReadLinesFromFile("day1") {
			if line == "" {
				elf := Elf{food: foods}
				elves <- elf
				foods = []Food{}
			} else if calories, err := strconv.Atoi(line); err == nil {
				foods = append(foods, Food{calories: calories})
			}
		}

		// This is to handle sending the very last item if there is no newline at the end of the file
		if len(foods) > 0 {
			elf := Elf{food: foods}
			elves <- elf
		}
		close(elves)
	}()
	return elves
}
