package goxai

import "github.com/D1CED/gox/gox"

// Score is a pseudo value rating a cell of a game. Higher values means more
// likelyhood to win.
type Score int

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
func unsidedFieldEval(b *gox.Board, rc [2]int) Score {
	if win, _ := b.CheckWin(); win {
		return Win
	}
	if b.Round() == gox.BoardSize {
		return Tie
	}
	var s Score
	switch rc {
	case [2]int{1, 1}:
		s = 2
	case [2]int{0, 0}, [2]int{2, 0}, [2]int{0, 2}, [2]int{2, 2}:
		s = 1
	case [2]int{0, 1}, [2]int{1, 0}, [2]int{1, 2}, [2]int{2, 1}:
		s = 0
	}
	return s
}

/*
func unsidedAdjacentEval(b *gox.Board, rc [2]int) Score {
*/
