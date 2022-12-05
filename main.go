package main

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/day1"
	"github.com/ryandem1/aoc_2022_go/day2"
	"github.com/ryandem1/aoc_2022_go/day3"
	"github.com/ryandem1/aoc_2022_go/day4"
	"log"
	"os"
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
	visualize := os.Args[2]

	fmt.Printf("\n--------------------------------------------------------%s--------------------------------------------------------\n\n", day)
	switch day {
	case "day1":
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
		if visualize == "-v" {
			day4.Visualize()
			break
		}
		printSolutions(
			solution{
				prompt: "In how many assignment pairs does one range fully contain the other?",
				answer: day4.Part1(),
			})
	default:
		log.Fatal("Unimplemented or invalid day!")
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
}
