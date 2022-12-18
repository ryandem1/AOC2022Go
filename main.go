package main

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/common"
	"github.com/ryandem1/aoc_2022_go/day1"
	"github.com/ryandem1/aoc_2022_go/day10"
	"github.com/ryandem1/aoc_2022_go/day2"
	"github.com/ryandem1/aoc_2022_go/day3"
	"github.com/ryandem1/aoc_2022_go/day4"
	"github.com/ryandem1/aoc_2022_go/day5"
	"github.com/ryandem1/aoc_2022_go/day6"
	"github.com/ryandem1/aoc_2022_go/day7"
	"github.com/ryandem1/aoc_2022_go/day8"
	"github.com/ryandem1/aoc_2022_go/day9"
	"log"
	"os"
	"strings"
)

// printSolutions will take a set of solution items that are part of the same day and print them in a standardized way
func printSolutions(solutions ...common.Solution) {
	for i, s := range solutions {
		fmt.Println(s.Prompt)
		fmt.Println(fmt.Sprintf("Part %d Answer:", i+1))
		fmt.Println(">>>", s.Answer, "<<<")
		fmt.Println()
	}
}

func main() {
	day := os.Args[1]
	visual := 0
	halfLine := strings.Repeat("-", 30)
	fullLine := strings.Repeat(halfLine, 2) + "----" // The extra ---- is to counter-balance the missing day txt

	if len(os.Args) == 3 {
		visual = 0
		parsedCount, err := fmt.Sscanf(os.Args[2], "-visual%d", &visual)
		if parsedCount == 0 || err != nil {
			panic("Invalid visualization parameter. Must be formatted like: '-visual1'")
		}
	}

	fmt.Printf("\n%s%s%s\n", halfLine, day, halfLine)
	switch day {
	case "day1":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day1.Part1(), day1.Part2())
	case "day2":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day2.Part1(), day2.Part2())
	case "day3":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day3.Part1(), day3.Part2())
	case "day4":
		availableVisualizations := 1

		switch visual {
		case 0: // No visualization option passed
			break
		case 1:
			fmt.Println("VISUAL 1: Camp Sections Cleaning Assignment Pairs")
			fmt.Println(fullLine)
			day4.Visualize()
		default:
			panic(fmt.Sprintf("Only have %d available visualization(s)", availableVisualizations))
		}
		if visual > 0 {
			break // Don't print solutions if visualize was passed
		}

		printSolutions(day4.Part1(), day4.Part2())
	case "day5":
		availableVisualizations := 2

		switch visual {
		case 0: // No visualization option passed
			break
		case 1:
			fmt.Println("VISUAL 1: CrateMover9000 Crane Actions Performed on Input")
			fmt.Println(fullLine)
			day5.Visualize(day5.CrateMover9000)
		case 2:
			fmt.Println("VISUAL 2: CrateMover9001 Crane Actions Performed on Input")
			fmt.Println(fullLine)
			day5.Visualize(day5.CrateMover9001)
		default:
			panic(fmt.Sprintf("Only have %d available visualization(s)", availableVisualizations))
		}
		if visual > 0 {
			break // Don't print solutions if visualize was passed
		}

		printSolutions(day5.Part1(), day5.Part2())
	case "day6":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day6.Day6(common.Part1), day6.Day6(common.Part2))
	case "day7":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day7.Part1(), day7.Part2())
	case "day8":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day8.Part1(), day8.Part2())
	case "day9":
		availableVisualizations := 2

		switch visual {
		case 0: // No visualization option passed
			break
		case 1:
			fmt.Println("VISUAL 1: Rope Movements per Motion (Rope Size: 2)")
			fmt.Println(fullLine)
			day9.VisualizeRopeMovements(2)
		case 2:
			fmt.Println("VISUAL 2: Rope Movements per Motion (Rope Size: 10)")
			fmt.Println(fullLine)
			day9.VisualizeRopeMovements(10)
		default:
			panic(fmt.Sprintf("Only have %d available visualization(s)", availableVisualizations))
		}
		if visual > 0 {
			break // Don't print solutions if visualize was passed
		}

		printSolutions(day9.Part1(), day9.Part2())
	case "day10":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(day10.Part1(), day10.Part2())
	default:
		log.Fatal("Unimplemented or invalid day!")
	}
	fmt.Println(fullLine)
}
