package ai

import (
	"testing"

	ttg "gox/gox"
)

type input struct {
	Game   *ttg.AIGame
	Symbol ttg.Symbol
	Field  [2]int
}

var testValue4 = ttg.AIGame{
	Board: [3][3]ttg.Symbol{
		{'O', 'O', 'X'}, {'X', 'O', 'X'}, {'O', 'X', 'X'},
	},
	Human:  'O',
	ArtInt: 'X',
}
var testValue5 = ttg.AIGame{
	Board: [3][3]ttg.Symbol{
		{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'O', 'X', 'O'},
	},
	Human:  'O',
	ArtInt: 'X',
}
var testValue6 = ttg.AIGame{'X', 'O', [3][3]ttg.Symbol{{'O', 0, 'X'}, {'X', 'X', 'O'}, {'O', 0, 'X'}}}

func TestValue(t *testing.T) {
	cases := []struct {
		in   input
		want Score
	}{
		{in: input{ttg.NewAIGame('O', 'X'), 'X', [2]int{1, 1}}, want: 2},
		{in: input{ttg.NewAIGame('O', 'X'), 'X', [2]int{0, 0}}, want: 1},
		{in: input{ttg.NewAIGame('O', 'X'), 'X', [2]int{1, 0}}, want: 0},
		{in: input{&testValue4, 'X', [2]int{1, 0}}, want: 10},
		{in: input{&testValue5, 'X', [2]int{1, 0}}, want: 0},
		{in: input{&testValue6, 'O', [2]int{0, 1}}, want: 0},
	}
	for idx, cas := range cases {
		got := sideBasedFieldEval(cas.in.Game, cas.in.Symbol, cas.in.Field)
		if got != cas.want {
			t.Errorf("Err in test %v: got %v, want %v",
				idx, got, cas.want)
		}
	}
}
