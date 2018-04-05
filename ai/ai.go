// Package ai provides a small API using the minimax algorithm with alpha-beta
// pruning.
package ai

import (
	"fmt"
	"time"

	"github.com/D1CED/gox/gox"
)

// difficulty type?

// Set sets the first field with the greatest score to AIGame.ArtInt.
func Set(g *gox.AIGame, difficulty int) error {
	r, c, _, err := EvalFields(g, difficulty)
	g.Board[r][c] = g.ArtInt
	return err
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
func EvalFields(g *gox.AIGame, difficulty int) (row, column int, scr Score, err error) {
	free := g.FreeFields()
	ch := make(chan struct {
		Score
		rc [2]int
		error
	}) //, len(free)
	for _, rc := range free {
		cp := *g // copy
		go func(rc [2]int, g *gox.AIGame, ch chan<- struct {
			Score
			rc [2]int
			error
		}) { // inline function-call
			scr, err := Evaluate(g, rc[0], rc[1], difficulty)
			select {
			case ch <- struct {
				Score
				rc [2]int
				error
			}{scr, rc, err}:
			case <-time.After(100 * time.Millisecond):
			}
		}(rc, &cp, ch)
	}
	max, maxIdx := minScore, [2]int{0, 0}
	for range free {
		v := <-ch
		scr, rc, err := v.Score, v.rc, v.error
		if err != nil {
			return 0, 0, 0, err
		}
		if scr == 10 {
			return rc[0], rc[1], 10, nil
		}
		if scr > max {
			max, maxIdx = scr, rc
		}
	}
	return maxIdx[0], maxIdx[1], max, nil
}
