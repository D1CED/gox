package main

import (
	"fmt"
	"os"
)

func modeAI(dfc int) {
	ans, err := InputLoop("Choose a side. X starts.", "X", "O")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hum := Symbol(ans[0])
	var aiSymb Symbol
	if hum == 'X' {
		aiSymb = 'O'
	} else {
		aiSymb = 'X'
	}
	g := &AIGame{Human: hum, ArtInt: aiSymb}

	for g.Round() < BoardSize {
		var cur Symbol
		if g.Round()%2 == 0 {
			cur = 'X'
		} else {
			cur = 'O'
		}
		if cur == hum {
			f, err := FieldInp(&g.Board)
			if err != nil {
				panic(err)
			}
			g.Board[f.row][f.col] = g.Human
		} else {
			err := Set(g, dfc)
			if err != nil {
				panic(err)
			}
		}
		if PrintWinDraw(&g.Board) {
			return
		}
	}
}
