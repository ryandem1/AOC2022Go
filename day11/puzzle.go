package day11

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
)

func Part1() (solution common.Solution) {
	solution.Prompt = `
Figure out which monkeys to chase by counting how many 
items they inspect over 20 rounds. 

What is the level of monkey business after 20 rounds 
of stuff-slinging simian shenanigans?
`
	numberOfItemsInspected := make(map[*keepAwayMonkey]int)

	// 20 rounds
	monkeys := readMonkeys()
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				numberOfItemsInspected[monkey]++

				monkey.operation(item)
				item.worryLevel = item.worryLevel / 3

				throwToId := monkey.test(item)
				throwToMonkey := common.FindOneObj(monkeys, func(monkey *keepAwayMonkey) bool {
					return monkey.id == throwToId
				})
				if throwToMonkey == nil {
					log.Panicf("Monkey with ID: %d not found!", throwToId)
				}

				monkey.throw(item, throwToMonkey)
			}
		}
	}

	highestInspectionCount := 0
	secondHighestInspectionCount := 0
	for _, inspectionCount := range numberOfItemsInspected {
		if inspectionCount > highestInspectionCount {
			secondHighestInspectionCount = highestInspectionCount
			highestInspectionCount = inspectionCount
		} else if inspectionCount > secondHighestInspectionCount {
			secondHighestInspectionCount = inspectionCount
		}
	}
	monkeyBusiness := highestInspectionCount * secondHighestInspectionCount

	solution.Answer = monkeyBusiness
	return solution
}

func Part2() (solution common.Solution) {
	solution.Prompt = ``
	return solution
}
