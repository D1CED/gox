package main

import (
	"flag"
	"fmt"
)

var (
	mode string
	dfc  int
)

// parses flags and picks gamemode.
func main() {
	flag.StringVar(&mode, "m", "AI", "Choose your gamemode 'AI' or 'MP'")
	flag.IntVar(&dfc, "d", 3, "Set difficulty for AI. 1 to 3.")
	flag.Parse()

	fmt.Println("Welcome to gox. A tic-tac-toe game.")

	switch mode {
	case "AI", "Ai", "ai", "A", "a":
		modeAI()
		return
	case "MP", "Mp", "mp", "M", "m":
		modeMP()
		return
	}
	fmt.Printf("No such mode '%s'. Quitting...\n", mode)
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
		r, c, err := FieldInp(&g.Board)
		if err != nil {
			panic(err)
		}
		g.Board[r][c] = s
		if PrintWinDraw(&g.Board) {
			return
		}
	}
}
