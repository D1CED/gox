package gox

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
