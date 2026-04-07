//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole1

import "errors"

// Game represents a TicTacToe game.
// This code is intentionally "smelly" for refactoring practice.
type Game struct {
	lastSymbol string
	board      *Board
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		lastSymbol: " ",
		board:      NewBoard(),
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
	//if the positions in first row are taken
	if g.board.TileAt(0, 0).Symbol != " " &&
		g.board.TileAt(0, 1).Symbol != " " &&
		g.board.TileAt(0, 2).Symbol != " " {
		//if first row is full with same symbol
		if g.board.TileAt(0, 0).Symbol == g.board.TileAt(0, 1).Symbol &&
			g.board.TileAt(0, 2).Symbol == g.board.TileAt(0, 1).Symbol {
			return g.board.TileAt(0, 0).Symbol
		}
	}

	//if the positions in second row are taken
	if g.board.TileAt(1, 0).Symbol != " " &&
		g.board.TileAt(1, 1).Symbol != " " &&
		g.board.TileAt(1, 2).Symbol != " " {
		//if middle row is full with same symbol
		if g.board.TileAt(1, 0).Symbol == g.board.TileAt(1, 1).Symbol &&
			g.board.TileAt(1, 2).Symbol == g.board.TileAt(1, 1).Symbol {
			return g.board.TileAt(1, 0).Symbol
		}
	}

	//if the positions in third row are taken
	if g.board.TileAt(2, 0).Symbol != " " &&
		g.board.TileAt(2, 1).Symbol != " " &&
		g.board.TileAt(2, 2).Symbol != " " {
		//if last row is full with same symbol
		if g.board.TileAt(2, 0).Symbol == g.board.TileAt(2, 1).Symbol &&
			g.board.TileAt(2, 2).Symbol == g.board.TileAt(2, 1).Symbol {
			return g.board.TileAt(2, 0).Symbol
		}
	}

	return " "
}

// Tile represents a position on the board.
type Tile struct {
	X      int
	Y      int
	Symbol string
}

// Board represents the game board.
type Board struct {
	plays []*Tile
}

// NewBoard creates a new board.
func NewBoard() *Board {
	b := &Board{plays: make([]*Tile, 0, 9)}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.plays = append(b.plays, &Tile{X: i, Y: j, Symbol: " "})
		}
	}
	return b
}

// TileAt returns the tile at the given position.
func (b *Board) TileAt(x, y int) *Tile {
	for _, t := range b.plays {
		if t.X == x && t.Y == y {
			return t
		}
	}
	return nil
}

// AddTileAt sets the symbol at the given position.
func (b *Board) AddTileAt(symbol string, x, y int) {
	for _, t := range b.plays {
		if t.X == x && t.Y == y {
			t.Symbol = symbol
			return
		}
	}
}
