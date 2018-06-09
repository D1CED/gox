package main

import "fmt"

// Score is a pseudo value rating a cell of a game. Higher values means more
// likelyhood to win.
type Score int

// FieldScore is a function that rates a filed of the board with a score
type FieldScore func(*Board, Field) Score

// Constants denoting common score values.
const (
	Win    Score = 10
	twoRow Score = 5
	oneRow Score = 2
	Tie    Score = 0
	Loss   Score = -10
)

func (s Score) String() string {
	switch s {
	case 10:
		return "Win"
	case -10:
		return "Loss"
	case 0:
		return "Tie"
	default:
		return fmt.Sprintf("%d", s)
	}
}

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

// aiPickFiled is a wrapper around alphabeta for easier usage.
func aiPickFiled(b *Board, rc Field, dfc int) Score {
	return alphabeta(b, rc, unsidedFieldEval, true, dfc*2, -200, 200)
}

// alphabeta is simmillar to negamax but with alpha-beta punning to reduce the
// amount of evaluated nodes.
func alphabeta(b *Board, rc Field, eval FieldScore, maximize bool, depth int,
	alpha, beta Score) Score {

	// steps++
	r, c := rc.row, rc.col
	symb := b[r][c]
	var opp Symbol
	if symb == X {
		opp = O
	} else {
		opp = X
	}

	if win, _ := b.CheckWin(); win || b.Round() == BoardSize ||
		depth == 0 {
		if !maximize {
			return -eval(b, rc)
		}
		return eval(b, rc)
	}

	scores := make([]Score, 0, len(b.FreeFields()))
	for _, rc := range b.FreeFields() {
		b[rc.row][rc.col] = opp
		s := -alphabeta(b, rc, eval, !maximize, depth-1, -beta, -alpha)
		alpha = max(alpha, s)
		scores = append(scores, s)
		b[rc.row][rc.col] = 0
		if alpha >= beta {
			break
		}
	}
	if maximize {
		return -max(scores...)
	}
	return max(scores...)
}

func max(i ...Score) Score {
	if len(i) == 0 {
		return 0
	}
	max := i[0]
	for _, v := range i {
		if v > max {
			max = v
		}
	}
	return max
}

// type node struct{ b *gox.Board; row, column int }
// var steps int // debug, dosen't work in parallel context

/*
// negamax is the negamax algorithm for rating fileds in a tic-tac-toe board
// for a given player. Provide a Board with rc already set.
func negamax(b *Board, rc [2]int, eval func(*Board, [2]int) Score,
	maximize bool, depth int) Score {

	// steps++
	r, c := rc[0], rc[1]
	symb := b[r][c]
	var opp Symbol
	if symb == 'X' {
		opp = 'O'
	} else {
		opp = 'X'
	}

	if win, _ := b.CheckWin(); win || b.Round() == BoardSize ||
		depth == 0 {
		if !maximize {
			return -eval(b, rc)
		}
		return eval(b, rc)
	}

	scores := make([]Score, 0, len(b.FreeFields()))
	for _, rc := range b.FreeFields() {
		b[rc[0]][rc[1]] = opp
		s := -negamax(b, rc, eval, !maximize, depth-1)
		scores = append(scores, s)
		b[rc[0]][rc[1]] = 0
	}
	if maximize {
		return -max(scores...)
	}
	return max(scores...)
}
*/
