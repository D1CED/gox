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

// NewMPGame is Game constructor.
func NewMPGame(p1, p2 Symbol) *MPGame {
	return &MPGame{Player1: p1, Player2: p2}
}

// NewAIGame is Game constructor.
func NewAIGame(hum, ai Symbol) *AIGame {
	return &AIGame{Human: hum, ArtInt: ai}
}
