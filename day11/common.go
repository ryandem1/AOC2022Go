package day11

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"log"
	"strconv"
	"strings"
)

// readMonkeys will read our input file and parse it out into keepAwayMonkey objects. Each monkey will contain some or
// no missingPackItem objects
func readMonkeys() []*keepAwayMonkey {
	var monkeys []*keepAwayMonkey

	for lines := range common.ReadChannelInChunks(common.ReadLinesFromFile("day11"), 7) {
		// region PARSE MONKEY ID
		monkeyId, err := strconv.Atoi(string(strings.Fields(lines[0])[1][0]))
		if err != nil {
			panic(err)
		}
		// endregion

		// region PARSE STARTING ITEMS
		var startingItems []*missingPackItem
		rawStartingItems := strings.Split(lines[1], ":")[1]
		for _, rawStartingItem := range strings.Split(rawStartingItems, ",") {
			worryLevel, err := strconv.Atoi(strings.TrimSpace(rawStartingItem))
			if err != nil {
				panic(err)
			}
			startingItems = append(startingItems, &missingPackItem{worryLevel: worryLevel})
		}
		// endregion

		// region PARSE OPERATION
		rawOperation := strings.Fields(strings.Split(lines[2], ":")[1])
		rawOperand1 := rawOperation[2]
		rawOperator := rawOperation[3]
		rawOperand2 := rawOperation[4]

		operation := func(item *missingPackItem) {
			var operand1 int
			var operand2 int

			if rawOperand1 == "old" {
				operand1 = item.worryLevel
			} else {
				operand1, err = strconv.Atoi(rawOperand1)
				if err != nil {
					panic(err)
				}
			}
			if rawOperand2 == "old" {
				operand2 = item.worryLevel
			} else {
				operand2, err = strconv.Atoi(rawOperand2)
				if err != nil {
					panic(err)
				}
			}

			switch rawOperator {
			case "+":
				item.worryLevel = operand1 + operand2
			case "*":
				item.worryLevel = operand1 * operand2
			default:
				log.Panicf("Unhandled operator: %s", rawOperator)
			}
		}
		// endregion

		// region PARSE TEST
		rawDivisibleBy := strings.Fields(lines[3])[len(strings.Fields(lines[3]))-1]
		divisibleBy, err := strconv.Atoi(rawDivisibleBy)
		if err != nil {
			panic(err)
		}

		trueTestThrowTarget, err := strconv.Atoi(strings.Fields(lines[4])[len(strings.Fields(lines[4]))-1])
		if err != nil {
			panic(err)
		}

		falseTestThrowTarget, err := strconv.Atoi(strings.Fields(lines[5])[len(strings.Fields(lines[5]))-1])
		if err != nil {
			panic(err)
		}

		test := func(item *missingPackItem) int {
			if item.worryLevel%divisibleBy == 0 {
				return trueTestThrowTarget
			} else {
				return falseTestThrowTarget
			}
		}
		// endregion

		monkeys = append(monkeys, &keepAwayMonkey{
			id:        monkeyId,
			items:     startingItems,
			operation: operation,
			test:      test,
		})

	}
	return monkeys
}

// a throw function that will throw a missingPackItem to a different keepAwayMonkey
func (monkey *keepAwayMonkey) throw(item *missingPackItem, target *keepAwayMonkey) {
	iItem := common.FindIndex(monkey.items, item)
	if iItem == -1 {
		panic("Monkey tried to throw an item that it doesn't have! Probably an implementation problem")
	}

	monkey.items = common.RemoveIndex(monkey.items, iItem)
	target.items = append(target.items, item)
}
