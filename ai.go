// Package goxai provides a small API using the minimax algorithm with
// alpha-beta pruning.

package main

import "fmt"

// difficulty type?

// Set sets the first field with the greatest score to AIGame.ArtInt.
func Set(g *AIGame, difficulty int) error {
	f, _ := EvalFields(g, difficulty)
	g.Board[f.row][f.col] = g.ArtInt
	return nil
}

// Evaluate returns a score from -10 to 10 for the given cell.
func Evaluate(g *AIGame, f Field, difficulty int) Score {
	if g.Board[f.row][f.col] != 0 {
		panic(fmt.Errorf("field (%v, %v) already set in Evaluate",
			f.row, f.col))
	}
	g.Board[f.row][f.col] = g.ArtInt
	scr := alphabeta(&g.Board, Field{f.row, f.col}, unsidedFieldEval,
		true, difficulty*2, -200, 200)
	g.Board[f.row][f.col] = 0
	return scr
}

// EvalFields returns the first cell with the greatest score and the score.
// It executes paralel.
func EvalFields(g *AIGame, difficulty int) (Field, Score) {

	type positionScore struct {
		s Score
		f Field
	}
	free := g.FreeFields()
	ch := make(chan positionScore) //, len(free)
	done := make(chan struct{})
	defer close(done)
	for i := range free {
		cp := *g // copy
		go func(f Field) {
			// difficulty from package-level scope
			scr := Evaluate(&cp, f, difficulty)
			select {
			case ch <- positionScore{scr, f}:
			case <-done:
			}
		}(free[i])
	}
	max, maxIdx := minScore, Field{0, 0}
	for range free {
		v := <-ch
		scr, f := v.s, v.f
		if scr == 10 {
			return f, 10
		}
		if scr > max {
			max, maxIdx = scr, f
		}
	}
	return maxIdx, max
}
