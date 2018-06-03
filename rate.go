package main

// Score is a pseudo value rating a cell of a game. Higher values means more
// likelyhood to win.
type Score int

// FieldScore is a function that rates a filed of the board with a score
type FieldScore func(*Board, Field) Score

// Constants denoting common score values.
const (
	maxScore = Score(^uint(0) >> 1)
	Win      = Score(10)
	twoRow   = Score(5)
	oneRow   = Score(2)
	Tie      = Score(0)
	Loss     = Score(-10)
	minScore = -(Score(^uint(0) >> 1)) - 1
)

// unsidedFieldEval rates a set field from 0 to 10.
func unsidedFieldEval(b *Board, rc Field) Score {
	if win, _ := b.CheckWin(); win {
		return Win
	}
	if b.Round() == BoardSize {
		return Tie
	}
	var s Score
	switch rc {
	case Field{1, 1}:
		s = 2
	case Field{0, 0}, Field{2, 0}, Field{0, 2}, Field{2, 2}:
		s = 1
	case Field{0, 1}, Field{1, 0}, Field{1, 2}, Field{2, 1}:
		s = 0
	}
	return s
}

/*
func unsidedAdjacentEval(b *gox.Board, rc [2]int) Score {
*/
