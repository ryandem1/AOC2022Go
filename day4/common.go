package day4

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
	"strings"
)

func readCleaningAssignmentPairs() chan [2]cleaningAssignment {
	pairs := make(chan [2]cleaningAssignment)

	go func() {
		for line := range common.ReadLinesFromFile("day4") {
			var pair [2]cleaningAssignment

			rawAssignments := strings.Split(line, ",")
			for i := 0; i < 2; i++ {
				startEnd := strings.Split(rawAssignments[i], "-")
				start, err := strconv.Atoi(startEnd[0])
				if err != nil {
					panic(err)
				}
				end, err := strconv.Atoi(startEnd[1])
				if err != nil {
					panic(err)
				}
				pair[i] = cleaningAssignment{
					start: start,
					end:   end,
				}
			}
			pairs <- pair
		}
		close(pairs)
	}()
	return pairs
}
