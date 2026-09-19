//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package smellytictactoe

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
