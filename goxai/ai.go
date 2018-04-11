// Package goxai provides a small API using the minimax algorithm with alpha-beta
// pruning.
package goxai

import (
	"fmt"
	"time"

	"github.com/D1CED/gox/gox"
)

// difficulty type?

// Set sets the first field with the greatest score to AIGame.ArtInt.
func Set(g *gox.AIGame, difficulty int) error {
	r, c, _, err := EvalFields(g, difficulty)
	if err != nil {
		return err
	}
	g.Board[r][c] = g.ArtInt
	return nil
}

// Evaluate returns a score from -10 to 10 for the given cell.
func Evaluate(g *gox.AIGame, row, column, difficulty int) (Score, error) {
	if g.Board[row][column] != 0 {
		return 0, fmt.Errorf("field (%v, %v) already set in Evaluate",
			row, column)
	}
	g.Board[row][column] = g.ArtInt
	scr := alphabeta(&g.Board, [2]int{row, column}, unsidedFieldEval,
		true, difficulty*2, -200, 200)
	g.Board[row][column] = 0
	return scr, nil
}

// EvalFields returns the first cell with the greatest score and the score.
// It executes paralel.
func EvalFields(g *gox.AIGame, difficulty int) (row, column int, scr Score,
	err error) {

	type positionScore struct {
		s  Score
		rc [2]int
		e  error
	}
	free := g.FreeFields()
	ch := make(chan positionScore) //, len(free)
	done := make(chan struct{})
	for i := range free {
		cp := *g // copy
		go func(rc [2]int) {
			// difficulty from package-level scope
			scr, err := Evaluate(&cp, rc[0], rc[1], difficulty)
			select {
			case ch <- positionScore{scr, rc, err}:
			case <-done:
			}
		}(free[i])
	}
	max, maxIdx := minScore, [2]int{0, 0}
	for range free {
		v := <-ch
		scr, rc, err := v.s, v.rc, v.e
		if err != nil {
			close(done)
			return 0, 0, 0, err
		}
		if scr == 10 {
			close(done)
			return rc[0], rc[1], 10, nil
		}
		if scr > max {
			max, maxIdx = scr, rc
		}
	}
	return maxIdx[0], maxIdx[1], max, nil
}
