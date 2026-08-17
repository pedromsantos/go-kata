//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole6

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
	if g.board.TileAt(x, y).Symbol != emptyPlay {
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
	for i := firstRow; i <= thirdRow; i++ {
		for j := firstColumn; j <= thirdColumn; j++ {
			b.plays = append(b.plays, &Tile{X: i, Y: j, Symbol: emptyPlay})
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

func (b *Board) findRowFullWithSamePlayer() string {
	if b.isRowFull(firstRow) && b.isRowFullWithSameSymbol(firstRow) {
		return b.TileAt(firstRow, firstColumn).Symbol
	}

	if b.isRowFull(secondRow) && b.isRowFullWithSameSymbol(secondRow) {
		return b.TileAt(secondRow, firstColumn).Symbol
	}

	if b.isRowFull(thirdRow) && b.isRowFullWithSameSymbol(thirdRow) {
		return b.TileAt(thirdRow, firstColumn).Symbol
	}

	return emptyPlay
}

func (b *Board) isRowFull(row int) bool {
	return b.TileAt(row, firstColumn).Symbol != emptyPlay &&
		b.TileAt(row, secondColumn).Symbol != emptyPlay &&
		b.TileAt(row, thirdColumn).Symbol != emptyPlay
}

func (b *Board) isRowFullWithSameSymbol(row int) bool {
	return b.TileAt(row, firstColumn).Symbol == b.TileAt(row, secondColumn).Symbol &&
		b.TileAt(row, thirdColumn).Symbol == b.TileAt(row, secondColumn).Symbol
}
