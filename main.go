package main

import (
	"fmt"
	"github.com/ryandem1/aoc_2022_go/day1"
	"github.com/ryandem1/aoc_2022_go/day2"
	"github.com/ryandem1/aoc_2022_go/day3"
	"log"
	"os"
)

func main() {
	day := os.Args[1]
	fmt.Printf("\n--------------------------------------------------------%s--------------------------------------------------------\n", day)
	switch day {
	case "day1":
		log.Println("Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?")
		log.Println("Part 1 Answer:", day1.Part1())

		log.Println(
			"Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying " +
				"in total?",
		)
		log.Println("Part 2 Answer:", day1.Part2())
	case "day2":
		log.Println("What would your total score be if everything goes exactly according to your strategy guide?")
		log.Println("Part 1 Answer:", day2.Part1())

		log.Println("What would your total score be if everything goes exactly according to your strategy guide?")
		log.Println("Part 2 Answer:", day2.Part2())
	case "day3":
		log.Println("What is the sum of the priorities of those item types?")
		log.Println("Part 1 Answer:", day3.Part1())
	default:
		log.Fatal("Unimplemented or invalid day!")
	}
	fmt.Println("\n--------------------------------------------------------------------------------------------------------------------")
}
