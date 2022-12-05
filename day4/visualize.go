package day4

import (
	"fmt"
	"strconv"
)

// Visualize will print out all the cleaning assignments as monospace textart
func Visualize() {
	for pair := range readCleaningAssignmentPairs() {
		fmt.Println()
		blankSpace := " -- "
		for _, assignment := range pair {
			line := ""

			for i := 0; i < assignment.start; i++ {
				line += blankSpace
			}
			for i := assignment.start; i < assignment.end; i++ {
				sectionId := strconv.Itoa(i)
				if i < 10 {
					sectionId = "0" + sectionId
				}
				line += " " + sectionId + " "
			}
			for i := assignment.end; i < TotalSections; i++ {
				line += blankSpace
			}
			fmt.Println(line)
		}
		fmt.Println()
	}
}
