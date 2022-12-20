package day11

// missingPackItem represents a single item that was in our pack as we were heading upstream that fell out. Each item
// has a corresponding worryLevel, which is an int representation of how worried we are to lose that item
type missingPackItem struct {
	worryLevel int
}

// keepAwayMonkey represents a single monkey participating in the game of keep-away for our items.
type keepAwayMonkey struct {
	startingItems []missingPackItem
}

// A keepAwayPlayer must have:
// an operation, which is a function of how our worryLevel for a missingPackItem will change upon inspection.
// a test is a function that a player will use to test our worryLevel for a missingPackItem and determine the next
// player to throw to
// a throw function that will throw a missingPackItem to a different keepAwayPlayer
type keepAwayPlayer interface {
	operation(int) int
	test(bool) *keepAwayPlayer
	throw(item *missingPackItem) *keepAwayPlayer
}
