//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole8

import "errors"

const (
	firstRow     = 0
	secondRow    = 1
	thirdRow     = 2
	firstColumn  = 0
	secondColumn = 1
	thirdColumn  = 2

	playerO  = "O"
	noPlayer = " "
)

// Game represents a TicTacToe game.
// This code is intentionally "smelly" for refactoring practice.
type Game struct {
	lastPlayer string
	board      *Board
}

// NewGame creates a new game.
func NewGame() *Game {
	return &Game{
		lastPlayer: noPlayer,
		board:      NewBoard(),
	}
}

// Play makes a move at the given position.
func (g *Game) Play(player string, x, y int) error {
	if err := g.validateFirstMove(player); err != nil {
		return err
	}
	if err := g.validatePlayer(player); err != nil {
		return err
	}
	if err := g.validatePositionIsEmpty(x, y); err != nil {
		return err
	}

	g.updateLastPlayer(player)
	g.updateBoard(player, x, y)
	return nil
}

func (g *Game) validateFirstMove(player string) error {
	if g.lastPlayer == noPlayer {
		if player == playerO {
			return errors.New("Invalid first player")
		}
	}
	return nil
}

func (g *Game) validatePlayer(player string) error {
	if player == g.lastPlayer {
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
	g.lastPlayer = player
}

func (g *Game) updateBoard(player string, x, y int) {
	g.board.AddTileAt(player, x, y)
}

// Winner returns the winning player or noPlayer if no winner.
func (g *Game) Winner() string {
	return g.board.findRowFullWithSamePlayer()
}

// Tile represents a position on the board and knows how to compare/update itself.
type Tile struct {
	x      int
	y      int
	player string
}

// NewTile creates a new tile.
func NewTile(x, y int, player string) *Tile {
	return &Tile{x: x, y: y, player: player}
}

// Player returns the tile's player.
func (t *Tile) Player() string {
	return t.player
}

// IsNotEmpty reports whether the tile has been played.
func (t *Tile) IsNotEmpty() bool {
	return t.Player() != noPlayer
}

// HasSamePlayerAs reports whether two tiles were played by the same player.
func (t *Tile) HasSamePlayerAs(other *Tile) bool {
	return t.Player() == other.Player()
}

// HasSameCoordinatesAs reports whether two tiles are at the same position.
func (t *Tile) HasSameCoordinatesAs(other *Tile) bool {
	return t.x == other.x && t.y == other.y
}

// UpdatePlayer sets the tile's player.
func (t *Tile) UpdatePlayer(newPlayer string) {
	t.player = newPlayer
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
			b.plays = append(b.plays, NewTile(x, y, noPlayer))
		}
	}
	return b
}

// TileAt returns the tile at the given position.
func (b *Board) TileAt(x, y int) *Tile {
	probe := NewTile(x, y, noPlayer)
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(probe) {
			return t
		}
	}
	return nil
}

// AddTileAt sets the player at the given position.
func (b *Board) AddTileAt(player string, x, y int) {
	tile := NewTile(x, y, player)
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(tile) {
			t.UpdatePlayer(player)
			return
		}
	}
}

func (b *Board) findRowFullWithSamePlayer() string {
	if b.isRowFull(firstRow) && b.isRowFullWithSamePlayer(firstRow) {
		return b.TileAt(firstRow, firstColumn).Player()
	}

	if b.isRowFull(secondRow) && b.isRowFullWithSamePlayer(secondRow) {
		return b.TileAt(secondRow, firstColumn).Player()
	}

	if b.isRowFull(thirdRow) && b.isRowFullWithSamePlayer(thirdRow) {
		return b.TileAt(thirdRow, firstColumn).Player()
	}

	return noPlayer
}

func (b *Board) isRowFull(row int) bool {
	return b.TileAt(row, firstColumn).IsNotEmpty() &&
		b.TileAt(row, secondColumn).IsNotEmpty() &&
		b.TileAt(row, thirdColumn).IsNotEmpty()
}

func (b *Board) isRowFullWithSamePlayer(row int) bool {
	return b.TileAt(row, firstColumn).HasSamePlayerAs(b.TileAt(row, secondColumn)) &&
		b.TileAt(row, thirdColumn).HasSamePlayerAs(b.TileAt(row, secondColumn))
}
