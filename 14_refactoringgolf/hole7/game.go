//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole7

import "errors"

const (
	firstRow     = 0
	secondRow    = 1
	thirdRow     = 2
	firstColumn  = 0
	secondColumn = 1
	thirdColumn  = 2

	playerO   = "O"
	emptyPlay = " "
)

// Game represents a TicTacToe game.
// This code is intentionally "smelly" for refactoring practice.
type Game struct {
	lastSymbol string
	board      *Board
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		lastSymbol: emptyPlay,
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
	if g.lastSymbol == emptyPlay {
		if player == playerO {
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
	if g.board.TileAt(x, y).IsNotEmpty() {
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
	return g.board.findRowFullWithSamePlayer()
}

// Tile represents a position on the board and knows how to compare/update itself.
type Tile struct {
	x      int
	y      int
	symbol string
}

// NewTile creates a new tile.
func NewTile(x, y int, symbol string) *Tile {
	return &Tile{x: x, y: y, symbol: symbol}
}

// Symbol returns the tile's symbol.
func (t *Tile) Symbol() string {
	return t.symbol
}

// IsNotEmpty reports whether the tile has been played.
func (t *Tile) IsNotEmpty() bool {
	return t.Symbol() != emptyPlay
}

// HasSameSymbolAs reports whether two tiles share the same symbol.
func (t *Tile) HasSameSymbolAs(other *Tile) bool {
	return t.Symbol() == other.Symbol()
}

// HasSameCoordinatesAs reports whether two tiles are at the same position.
func (t *Tile) HasSameCoordinatesAs(other *Tile) bool {
	return t.x == other.x && t.y == other.y
}

// UpdateSymbol sets the tile's symbol.
func (t *Tile) UpdateSymbol(newSymbol string) {
	t.symbol = newSymbol
}

// Board represents the game board.
type Board struct {
	plays []*Tile
}

// NewBoard creates a new board.
func NewBoard() *Board {
	b := &Board{plays: make([]*Tile, 0, 9)}
	for x := firstRow; x <= thirdRow; x++ {
		for y := firstColumn; y <= thirdColumn; y++ {
			b.plays = append(b.plays, NewTile(x, y, emptyPlay))
		}
	}
	return b
}

// TileAt returns the tile at the given position.
func (b *Board) TileAt(x, y int) *Tile {
	probe := NewTile(x, y, emptyPlay)
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(probe) {
			return t
		}
	}
	return nil
}

// AddTileAt sets the symbol at the given position.
func (b *Board) AddTileAt(symbol string, x, y int) {
	probe := NewTile(x, y, symbol)
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(probe) {
			t.UpdateSymbol(symbol)
			return
		}
	}
}

func (b *Board) findRowFullWithSamePlayer() string {
	if b.isRowFull(firstRow) && b.isRowFullWithSameSymbol(firstRow) {
		return b.TileAt(firstRow, firstColumn).Symbol()
	}

	if b.isRowFull(secondRow) && b.isRowFullWithSameSymbol(secondRow) {
		return b.TileAt(secondRow, firstColumn).Symbol()
	}

	if b.isRowFull(thirdRow) && b.isRowFullWithSameSymbol(thirdRow) {
		return b.TileAt(thirdRow, firstColumn).Symbol()
	}

	return emptyPlay
}

func (b *Board) isRowFull(row int) bool {
	return b.TileAt(row, firstColumn).IsNotEmpty() &&
		b.TileAt(row, secondColumn).IsNotEmpty() &&
		b.TileAt(row, thirdColumn).IsNotEmpty()
}

func (b *Board) isRowFullWithSameSymbol(row int) bool {
	return b.TileAt(row, firstColumn).HasSameSymbolAs(b.TileAt(row, secondColumn)) &&
		b.TileAt(row, thirdColumn).HasSameSymbolAs(b.TileAt(row, secondColumn))
}
