package main

import (
	"flag"
	"fmt"
)

func setup() (mode string, difficulty int) {
	m := flag.String("m", "AI", "Choose your gamemode 'AI' or 'MP'")
	dfc := flag.Int("d", 3, "Set difficulty for AI. 1 to 3.")
	flag.Parse()
	if *dfc < 0 {
		*dfc = 0
	}
	if *dfc > 3 {
		*dfc = 3
	}
	return *m, *dfc
}

// parses flags and picks gamemode.
func main() {
	mode, dfc := setup()
	fmt.Println("Welcome to gox. A tic-tac-toe game.")

	switch mode {
	case "AI", "Ai", "ai", "A", "a":
		modeAI(dfc)
	case "MP", "Mp", "mp", "M", "m":
		modeMP()
	default:
		fmt.Printf("No such mode '%s'. Quitting...\n", mode)
	}
}

func modeMP() {
	g := &MPGame{Player1: 'X', Player2: 'O'}
	// max 9 rounds or someone wins
	for g.Round() < BoardSize {
		var s Symbol
		if g.Round()%2 == 0 {
			s = g.Player1
		} else {
			s = g.Player2
		}
		f, err := FieldInp(&g.Board)
		if err != nil {
			panic(err)
		}
		g.Board[f.row][f.col] = s
		if PrintWinDraw(&g.Board) {
			break
		}
	}
}
