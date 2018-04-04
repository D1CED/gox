package main

import (
	"gox/ai"
	"gox/gox"
	"gox/utils"
)

func modeAI() {
	if dfc < 0 {
		dfc = 0
	}
	if dfc > 3 {
		dfc = 3
	}
	ans, err := utils.InputLoop("Choose a side. X starts.", "X", "O")
	if err != nil {
		panic(err)
	}
	hum := gox.Symbol(ans[0])
	var aiSymb gox.Symbol
	if hum == 'X' {
		aiSymb = 'O'
	} else {
		aiSymb = 'X'
	}
	g := gox.NewAIGame(hum, aiSymb)

	for g.Round() < gox.BoardSize {
		var cur gox.Symbol
		if g.Round()%2 == 0 {
			cur = 'X'
		} else {
			cur = 'O'
		}
		if cur == hum {
			r, c, err := utils.FieldInp(&g.Board)
			if err != nil {
				panic(err)
			}
			g.Board[r][c] = g.Human
		} else {
			err := ai.Set(g, dfc)
			if err != nil {
				panic(err)
			}
		}
		if utils.PrintWinDraw(&g.Board) {
			return
		}
	}
}