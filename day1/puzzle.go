package day1

import "github.com/ryandem1/aoc_2022_go/common"

func Part1() (solution common.Solution) {
	solution.Prompt = `
Find the Elf carrying the most Calories.
How many totalCalories is that Elf carrying?
`
	var totalCalories int

	for elf := range readElves() {
		calories := elf.totalCaloriesHeld()
		if calories > totalCalories {
			totalCalories = calories
		}
	}
	solution.Answer = totalCalories
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = `
Find the Elf carrying the most Calories.
How many totalCalories is that Elf carrying?
`
	var totalCalories int

	topThreeTotalCalories := [3]int{}

	for elf := range readElves() {
		calories := elf.totalCaloriesHeld()

		for i, topCalories := range topThreeTotalCalories {
			if calories > topCalories {
				j := i
				for j < 2 {
					topThreeTotalCalories[i+1] = topThreeTotalCalories[i]
					j++
				}
				topThreeTotalCalories[i] = calories
				break
			}
		}
	}

	for _, topCalories := range topThreeTotalCalories {
		totalCalories += topCalories
	}

	solution.Answer = totalCalories
	return solution
}
