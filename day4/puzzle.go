package day4

// Part1 In how many assignment pairs does one range fully contain the other?
func Part1() (fullyOverlappingPairsCount int) {
	for pair := range readCleaningAssignmentPairs() {
		for i := 0; i < 2; i++ {
			if pair[i].start <= pair[i^1].start && pair[i].end >= pair[i^1].end {
				fullyOverlappingPairsCount++
				break
			}
		}
	}
	return fullyOverlappingPairsCount
}

// Part2 In how many assignment pairs do the ranges overlap?
func Part2() (overlappingPairsCount int) {
	return overlappingPairsCount
}
