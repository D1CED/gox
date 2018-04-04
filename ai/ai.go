// Package ai provides a small API using the minimax algorithm with alpha-beta
// pruning.
package ai

import (
	"fmt"

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
		return 0, fmt.Errorf("Filed (%v, %v) already set in Evaluate.",
			row, column)
	}
	g.Board[row][column] = g.ArtInt
	scr := alphabeta(&g.Board, [2]int{row, column}, unsidedFieldEval,
		true, difficulty*2, -200, 200)
	g.Board[row][column] = 0
	return scr, nil
}

// EvalFields returns the first cell with the greatest score and the score.
func EvalFields(g *gox.AIGame, difficulty int) (row, column int, scr Score, err error) {
	free := g.FreeFields()
	scores := make([]Score, len(free))
	for i, rc := range free {
		scr, err := Evaluate(g, rc[0], rc[1], difficulty)
		scores[i] = scr
		if scr == 10 {
			break
		}
		if err != nil {
			return 0, 0, 0, err
		}
	}
	max, maxIdx := minScore, -1
	for i, v := range scores {
		if v > max {
			maxIdx, max = i, v
		}
	}
	fmt.Printf("steps: %v; heuristics: %v\n", steps, scores) // debug
	steps = 0
	return free[maxIdx][0], free[maxIdx][1], max, nil
}
