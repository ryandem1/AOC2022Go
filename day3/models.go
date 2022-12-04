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
