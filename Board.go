package main

import "fmt"

// Board represents a tic-tac-toe field. (0, 0) is upper left corner, (2, 2) is
// lower right.
type Board [3][3]Symbol

// Field represents one filed of the Board.
type Field struct{ row, col int }

const BoardSize = 3 * 3

// String returns you the game board as multi-line string. Preferably to print.
func (b *Board) String() string {
	return fmt.Sprintf(`
    1   2   3
  +---+---+---+
a | %s | %s | %s |
  +---+---+---+
b | %s | %s | %s |
  +---+---+---+
c | %s | %s | %s |
  +---+---+---+
`,
		b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2],
		b[2][0], b[2][1], b[2][2],
	)
}

// CheckWin returns true if a player won and returns the corresponding symbol.
func (b *Board) CheckWin() (win bool, player Symbol) {
	// vertical and horizontal check
	for i := 0; i < 3; i++ {
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] && b[i][0] != 0 {
			return true, b[i][0]
		}
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] && b[0][i] != 0 {
			return true, b[0][i]
		}
	}
	// diagonal check
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[0][0] != 0 {
		return true, b[0][0]
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[0][2] != 0 {
		return true, b[0][2]
	}
	return
}

// Round returns the number of occupied fields.
func (b *Board) Round() int {
	ctr := 0
	for i := range b {
		for _, val := range b[i] {
			if val != 0 {
				ctr++
			}
		}
	}
	return ctr
}

// FreeFields returns a slice containig (row, col) of unset fields.
func (b *Board) FreeFields() []Field {
	free := make([]Field, 0, 9)
	for idx := range b {
		for jdx, elm := range b[idx] {
			if elm == 0 {
				free = append(free, Field{idx, jdx})
			}
		}
	}
	return free
}

// MPGame represents a tic-tac-toe game with a board and two players.
type MPGame struct {
	Player1 Symbol
	Player2 Symbol
	Board
}

// AIGame represents a tic-tac-toe game with one player being human, the other
// being an ai.
type AIGame struct {
	Human  Symbol
	ArtInt Symbol
	Board
}

// Symbol holds a character representing a player. Make sure they are uniqe and non-zero.
type Symbol byte

const X, O = 'X', 'O'

func (s Symbol) String() string {
	if s == 0 {
		return " "
	}
	return fmt.Sprintf("%c", s)
}
