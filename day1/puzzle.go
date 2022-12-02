package day1

// Part1 Find the Elf carrying the most Calories. How many totalCalories is that Elf carrying?
func Part1() (totalCalories int) {
	for elf := range readElves() {
		calories := elf.totalCaloriesHeld()
		if calories > totalCalories {
			totalCalories = calories
		}
	}
	return totalCalories
}

// Part2 Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func Part2() (totalCalories int) {
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

	return totalCalories
}
