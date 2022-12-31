package day11

// missingPackItem represents a single item that was in our pack as we were heading upstream that fell out. Each item
// has a corresponding worryLevel, which is an int representation of how worried we are to lose that item
type missingPackItem struct {
	worryLevel int
}

// keepAwayMonkey represents a single monkey participating in the game of keep-away for our items.
// an id is a numerical id that identifies a unique monkey
// an operation, which is a function of how our worryLevel for a missingPackItem will change upon inspection.
// a test is a function that a player will use to test our worryLevel for a missingPackItem and determine the next
// player to throw to. Returns the id of the monkey that the item should be thrown to
type keepAwayMonkey struct {
	id        int
	items     []*missingPackItem
	operation func(item *missingPackItem)
	test      func(item *missingPackItem) int
}
