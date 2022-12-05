package day4

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
