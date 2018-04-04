package utils

import (
	"fmt"
	"os"

	"github.com/D1CED/gox/gox"
)

// PrintWinDraw returns true if win or draw happend and prints this on screen.
func PrintWinDraw(b *gox.Board) bool {
	if win, symb := b.CheckWin(); win {
		fmt.Printf("Player %c wins. Congrats!\n", symb)
		fmt.Println(b)
		return true
	}
	if b.Round() == gox.BoardSize {
		fmt.Println("Oh, a draw.")
		fmt.Println(b)
		return true
	}
	return false
}

// InputLoop prints a message and gives the user a choice from a string.
// If the user enters  "q" it interrups the program.
func InputLoop(msg string, s ...string) (string, error) {
	for {
		var inp string
		fmt.Println(msg)
		fmt.Printf("Enter one of the following expressions: %v.\n", s)
		fmt.Println("Or hit 'q' to quit.")
		fmt.Print(">>> ")
		_, err := fmt.Scanln(&inp)
		if err != nil {
			return "", err
		}
		if inp == "q" {
			fmt.Println("Quitting...")
			os.Exit(0)
		}
		for _, v := range s {
			if inp == v {
				return v, nil
			}
		}
		fmt.Println("Sry invalid Input. Try again.")
	}
}

// FieldInp blaah with low level shit hacks
func FieldInp(b *gox.Board) (row, col int, err error) {
	free := b.FreeFields()
	choose := make([]string, len(free))
	for i, rc := range free {
		choose[i] = fmt.Sprintf("%c%c", rc[0]+97, rc[1]+49)
	}
	choice, err := InputLoop(b.String(), choose...)
	if err != nil {
		return
	}
	return int(byte(choice[0]) - 97), int(byte(choice[1]) - 49), nil
}
