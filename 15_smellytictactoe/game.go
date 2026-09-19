//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package smellytictactoe

import "errors"

// Game represents a TicTacToe game.
// This code is intentionally "smelly" for refactoring practice.
// Your goal is to identify and fix the code smells while maintaining functionality.
type Game struct {
	lastSymbol            string
	board                 *Board
	rowWinnerChecker      *RowWinnerChecker
	columnWinnerChecker   *ColumnWinnerChecker
	diagonalWinnerChecker *DiagonalWinnerChecker
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		lastSymbol:            " ",
		board:                 NewBoard(),
		rowWinnerChecker:      &RowWinnerChecker{},
		columnWinnerChecker:   &ColumnWinnerChecker{},
		diagonalWinnerChecker: &DiagonalWinnerChecker{},
	}
}

// Play makes a move at the given position.
func (g *Game) Play(symbol string, x, y int) error {
	//if first move
	if g.lastSymbol == " " {
		//if player is X
		if symbol == "O" {
			return errors.New("Invalid first player")
		}
	} else if symbol == g.lastSymbol {
		//if not first move but player repeated
		return errors.New("Invalid next player")
	} else if g.board.TileAt(x, y).Symbol != " " {
		//if not first move but play on an already played tile
		return errors.New("Invalid position")
	}

	// update game state
	g.lastSymbol = symbol
	g.board.AddTileAt(symbol, x, y)
	return nil
}

// Winner returns the winning symbol or " " if no winner.
func (g *Game) Winner() string {
	rowWinner := g.rowWinnerChecker.Check(g.board)
	if rowWinner != " " {
		return rowWinner
	}

	columnWinner := g.columnWinnerChecker.Check(g.board)
	if columnWinner != " " {
		return columnWinner
	}

	return g.diagonalWinnerChecker.Check(g.board)
}
