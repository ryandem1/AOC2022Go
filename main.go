package main

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/day1"
	"github.com/ryandem1/aoc_2022_go/day2"
	"github.com/ryandem1/aoc_2022_go/day3"
	"github.com/ryandem1/aoc_2022_go/day4"
	"github.com/ryandem1/aoc_2022_go/day5"
	"log"
	"os"
	"strings"
)

type solution struct {
	prompt string
	answer any
}

// printSolutions will take a set of solution items that are part of the same day and print them in a standardized way
func printSolutions(solutions ...solution) {
	for i, s := range solutions {
		log.Println(s.prompt)
		log.Println(fmt.Sprintf("Part %d Answer:", i+1))
		log.Println(">>>", s.answer, "<<<")
		fmt.Println()
	}
}

func main() {
	day := os.Args[1]
	visual := 0
	halfLine := strings.Repeat("-", 56)
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
		printSolutions(
			solution{
				prompt: "Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?",
				answer: day1.Part1(),
			},
			solution{
				prompt: "Find the top three Elves carrying the most Calories. How many Calories are those Elves " +
					"carrying in total?",
				answer: day1.Part2(),
			})
	case "day2":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(
			solution{
				prompt: "What would your total score be if everything goes exactly according to your strategy guide?",
				answer: day2.Part1(),
			},
			solution{
				prompt: "What would your total score be if everything goes exactly according to your strategy guide?",
				answer: day2.Part2(),
			})
	case "day3":
		if visual != 0 {
			panic(fmt.Sprintf("No visualization for %s", day))
		}
		printSolutions(
			solution{
				prompt: "What is the sum of the priorities of those item types?",
				answer: day3.Part1(),
			},
			solution{
				prompt: "What is the sum of the priorities of those item types?",
				answer: day3.Part2(),
			})
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

		printSolutions(
			solution{
				prompt: "In how many assignment pairs does one range fully contain the other?",
				answer: day4.Part1(),
			},
			solution{
				prompt: "In how many assignment pairs do the ranges overlap?",
				answer: day4.Part2(),
			})
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

		printSolutions(
			solution{
				prompt: "After the rearrangement procedure completes, what crate ends up on top of each stack?",
				answer: day5.Part1(),
			},
			solution{
				prompt: "After the rearrangement procedure completes, what crate ends up on top of each stack?",
				answer: day5.Part2(),
			})

	default:
		log.Fatal("Unimplemented or invalid day!")
	}
	fmt.Println(fullLine)
}
