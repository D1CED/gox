package main

func modeAI() {
	if dfc < 0 {
		dfc = 0
	}
	if dfc > 3 {
		dfc = 3
	}
	ans, err := InputLoop("Choose a side. X starts.", "X", "O")
	if err != nil {
		panic(err)
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
			r, c, err := FieldInp(&g.Board)
			if err != nil {
				panic(err)
			}
			g.Board[r][c] = g.Human
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
