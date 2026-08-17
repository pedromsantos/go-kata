//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole2

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
	if err := g.validateFirstMove(symbol); err != nil {
		return err
	}
	if err := g.validatePlayer(symbol); err != nil {
		return err
	}
	if err := g.validatePositionIsEmpty(x, y); err != nil {
		return err
	}

	g.updateLastPlayer(symbol)
	g.updateBoard(symbol, x, y)
	return nil
}

func (g *Game) validateFirstMove(player string) error {
	if g.lastSymbol == " " {
		if player == "O" {
			return errors.New("Invalid first player")
		}
	}
	return nil
}

func (g *Game) validatePlayer(player string) error {
	if player == g.lastSymbol {
		return errors.New("Invalid next player")
	}
	return nil
}

func (g *Game) validatePositionIsEmpty(x, y int) error {
	if g.board.TileAt(x, y).Symbol != " " {
		return errors.New("Invalid position")
	}
	return nil
}

func (g *Game) updateLastPlayer(player string) {
	g.lastSymbol = player
}

func (g *Game) updateBoard(player string, x, y int) {
	g.board.AddTileAt(player, x, y)
}

// Winner returns the winning symbol or " " if no winner.
func (g *Game) Winner() string {
	if g.isFirstRowFull() && g.isFirstRowFullWithSameSymbol() {
		return g.board.TileAt(0, 0).Symbol
	}

	if g.isSecondRowFull() && g.isSecondRowFullWithSameSymbol() {
		return g.board.TileAt(1, 0).Symbol
	}

	if g.isThirdRowFull() && g.isThirdRowFullWithSameSymbol() {
		return g.board.TileAt(2, 0).Symbol
	}

	return " "
}

func (g *Game) isFirstRowFull() bool {
	return g.board.TileAt(0, 0).Symbol != " " &&
		g.board.TileAt(0, 1).Symbol != " " &&
		g.board.TileAt(0, 2).Symbol != " "
}

func (g *Game) isFirstRowFullWithSameSymbol() bool {
	return g.board.TileAt(0, 0).Symbol == g.board.TileAt(0, 1).Symbol &&
		g.board.TileAt(0, 2).Symbol == g.board.TileAt(0, 1).Symbol
}

func (g *Game) isSecondRowFull() bool {
	return g.board.TileAt(1, 0).Symbol != " " &&
		g.board.TileAt(1, 1).Symbol != " " &&
		g.board.TileAt(1, 2).Symbol != " "
}

func (g *Game) isSecondRowFullWithSameSymbol() bool {
	return g.board.TileAt(1, 0).Symbol == g.board.TileAt(1, 1).Symbol &&
		g.board.TileAt(1, 2).Symbol == g.board.TileAt(1, 1).Symbol
}

func (g *Game) isThirdRowFull() bool {
	return g.board.TileAt(2, 0).Symbol != " " &&
		g.board.TileAt(2, 1).Symbol != " " &&
		g.board.TileAt(2, 2).Symbol != " "
}

func (g *Game) isThirdRowFullWithSameSymbol() bool {
	return g.board.TileAt(2, 0).Symbol == g.board.TileAt(2, 1).Symbol &&
		g.board.TileAt(2, 2).Symbol == g.board.TileAt(2, 1).Symbol
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
	tile := &Tile{X: x, Y: y, Symbol: symbol}
	_ = tile

	for _, t := range b.plays {
		if t.X == x && t.Y == y {
			t.Symbol = symbol
			return
		}
	}
}
