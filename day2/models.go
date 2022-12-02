package day2

type handShape int32

const (
	rock handShape = iota
	paper
	scissors
)

type player struct {
	score int
}

// newPlayer will create a new player, by default a player is not a winner
func newPlayer() *player {
	obj := player{score: 0}
	return &obj
}

type choice struct {
	player *player
	shape  handShape
}

type round struct {
	choice1 *choice
	choice2 *choice
	winner  *player
}

func newRound(choice1 *choice, choice2 *choice) *round {
	obj := &round{
		choice1: choice1,
		choice2: choice2,
		winner:  nil,
	}
	return obj
}

type game struct {
	player1 *player
	player2 *player
	rounds  []*round
}

// newGame will create a new game from pointers to players. Games start with no rounds, to add rounds, see addRound
func newGame(player1 *player, player2 *player) *game {
	game := game{
		player1: player1,
		player2: player2,
		rounds:  nil,
	}
	return &game
}

// addRound will add a new round to a game's rounds
func (game *game) addRound(player1Shape handShape, player2Shape handShape) {
	choice1 := &choice{
		player: game.player1,
		shape:  player1Shape,
	}
	choice2 := &choice{
		player: game.player2,
		shape:  player2Shape,
	}
	round := newRound(choice1, choice2)
	game.rounds = append(game.rounds, round)
}
