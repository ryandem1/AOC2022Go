package day3

// elfRucksack objects contain contents. Each rucksack has two large compartments. All items of a given type are meant
// to go into exactly one of the two compartments.
type elfRucksack struct {
	compartment1 string
	compartment2 string
}

// newRucksack will take a contents string (example: vJrwpWtwJgWrhcsFMMfFFhFp) and turn it into an elfRucksack obj
func newRucksack(contents string) elfRucksack {
	contentLength := len(contents)
	rucksack := elfRucksack{
		compartment1: contents[:contentLength/2],
		compartment2: contents[contentLength/2:],
	}
	return rucksack
}

// getPriorities will return the map of priorities that each item has
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func getPriorities() map[rune]int {
	priorities := make(map[rune]int)
	pValue := 0 // priority value

	for letter := 'a'; letter <= 'z'; letter++ {
		pValue++
		priorities[letter] = pValue
	}

	for letter := 'A'; letter <= 'Z'; letter++ {
		pValue++
		priorities[letter] = pValue
	}
	return priorities
}
