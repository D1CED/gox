package main

// type node struct{ b *gox.Board; row, column int }
// var steps int // debug, dosen't work in parallel context

// negamax is the negamax algorithm for rating fileds in a tic-tac-toe board
// for a given player. Provide a Board with rc already set.
func negamax(b *Board, rc [2]int, eval func(*Board, [2]int) Score,
	maximize bool, depth int) Score {

	// steps++
	r, c := rc[0], rc[1]
	symb := b[r][c]
	var opp Symbol
	if symb == 'X' {
		opp = 'O'
	} else {
		opp = 'X'
	}

	if win, _ := b.CheckWin(); win || b.Round() == BoardSize ||
		depth == 0 {
		if !maximize {
			return -eval(b, rc)
		}
		return eval(b, rc)
	}

	scores := make([]Score, 0, len(b.FreeFields()))
	for _, rc := range b.FreeFields() {
		b[rc[0]][rc[1]] = opp
		s := -negamax(b, rc, eval, !maximize, depth-1)
		scores = append(scores, s)
		b[rc[0]][rc[1]] = 0
	}
	if maximize {
		return -max(scores...)
	}
	return max(scores...)
}

// alphabeta is simmillar to negamax but with alpha-beta breaks to reduce the
// amount of evaluated nodes.
func alphabeta(b *Board, rc [2]int, eval func(*Board, [2]int) Score,
	maximize bool, depth int, alpha, beta Score) Score {

	// steps++
	r, c := rc[0], rc[1]
	symb := b[r][c]
	var opp Symbol
	if symb == 'X' {
		opp = 'O'
	} else {
		opp = 'X'
	}

	if win, _ := b.CheckWin(); win || b.Round() == BoardSize ||
		depth == 0 {
		if !maximize {
			return -eval(b, rc)
		}
		return eval(b, rc)
	}

	scores := make([]Score, 0, len(b.FreeFields()))
	for _, rc := range b.FreeFields() {
		b[rc[0]][rc[1]] = opp
		s := -alphabeta(b, rc, eval, !maximize, depth-1, -beta, -alpha)
		alpha = max(alpha, s)
		scores = append(scores, s)
		b[rc[0]][rc[1]] = 0
		if alpha >= beta {
			break
		}
	}
	if maximize {
		return -max(scores...)
	}
	return max(scores...)
}

func max(i ...Score) Score {
	max := minScore
	for _, v := range i {
		if v > max {
			max = v
		}
	}
	return max
}
