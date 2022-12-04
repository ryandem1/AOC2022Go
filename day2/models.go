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
