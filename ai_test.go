package main

import "testing"

func TestPrimFieldEval(t *testing.T) {
	var tests = []struct {
		in    *AIGame
		field Field
		want  Score
	}{
		{
			// ArtInt win
			in: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{O, O, X}, {X, O, X}, {O, X, X}}},
			field: Field{0, 2},
			want:  Win,
		}, {
			// Tie
			in: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{X, O, X}, {X, O, X}, {O, X, O}}},
			field: Field{1, 1},
			want:  Tie,
		}, {
			in: &AIGame{
				Human:  X,
				ArtInt: O,
				Board:  Board{{O, 0, X}, {X, X, O}, {O, 0, X}}},
			field: Field{2, 2},
			want:  1,
		},
	}
	for idx, test := range tests {
		got := unsidedFieldEval(&test.in.Board, test.field)
		if got != test.want {
			t.Errorf("Err in test %v: got %v, want %v",
				idx, got, test.want)
		}
	}
}

func TestAIEval(t *testing.T) {
	var tests = []struct {
		in    *AIGame
		field Field
		dfc   int
		want  Score
	}{
		{
			// ArtInt win
			in: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{O, 0, 0}, {X, O, X}, {O, X, X}}},
			field: Field{0, 2},
			dfc:   3,
			want:  Win,
		}, {
			// Tie
			in: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{X, O, X}, {X, O, 0}, {O, X, O}}},
			field: Field{0, 2},
			dfc:   3,
			want:  Tie,
		},
	}
	for idx, test := range tests {
		got := Evaluate(test.in, test.field, test.dfc)
		if got != test.want {
			t.Errorf("Err in test %v: got %v, want %v",
				idx, got, test.want)
		}
	}
}

func TestEvalField(t *testing.T) {
	var tests = []struct {
		game  *AIGame
		dfc   int
		field Field
		score Score
	}{
		{
			// ArtInt win
			game: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{O, O, X}, {X, O, X}, {O, X, X}}},
			dfc:   3,
			field: Field{0, 2},
			score: Win,
		}, {
			// Tie
			game: &AIGame{
				Human:  O,
				ArtInt: X,
				Board:  Board{{X, O, X}, {X, O, X}, {O, X, O}}},
			dfc:   3,
			field: Field{1, 1},
			score: Tie,
		}, {
			game: &AIGame{
				Human:  X,
				ArtInt: O,
				Board:  Board{{O, 0, X}, {X, X, O}, {O, 0, X}}},
			dfc:   3,
			field: Field{2, 1},
			score: Win,
		},
	}
	for idx, test := range tests {
		f, s := EvalFields(test.game, test.dfc)
		if f != test.field || s != test.score {
			t.Errorf("Err in test %v: got %v, want %v",
				idx, f, test.field)
		}
	}
}
