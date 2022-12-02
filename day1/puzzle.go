package day1

// Part1 Find the Elf carrying the most Calories. How many totalCalories is that Elf carrying?
func Part1() (totalCalories int) {
	for elf := range readElves() {
		caloriesHeld := 0
		for _, food := range elf.food {
			caloriesHeld += food.calories
		}
		if caloriesHeld > totalCalories {
			totalCalories = caloriesHeld
		}
	}
	return totalCalories
}
