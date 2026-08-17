//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole11

import "errors"

// Row is a board row, 0 through 2.
type Row int

// Column is a board column, 0 through 2.
type Column int

// Player is a board symbol: noPlayer, "X" or "O".
type Player string

const (
	firstRow  Row = 0
	secondRow Row = 1
	thirdRow  Row = 2

	firstColumn  Column = 0
	secondColumn Column = 1
	thirdColumn  Column = 2

	playerO  Player = "O"
	noPlayer Player = " "
)

// Game represents a TicTacToe game.
// This code is intentionally "smelly" for refactoring practice.
type Game struct {
	lastPlayer Player
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
func (g *Game) Play(player Player, x Row, y Column) error {
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
	g.updateBoard(NewTile(x, y, player))
	return nil
}

func (g *Game) validateFirstMove(player Player) error {
	if g.lastPlayer == noPlayer {
		if player == playerO {
			return errors.New("Invalid first player")
		}
	}
	return nil
}

func (g *Game) validatePlayer(player Player) error {
	if player == g.lastPlayer {
		return errors.New("Invalid next player")
	}
	return nil
}

func (g *Game) validatePositionIsEmpty(x Row, y Column) error {
	if g.board.isTilePlayedAt(x, y) {
		return errors.New("Invalid position")
	}
	return nil
}

func (g *Game) updateLastPlayer(player Player) {
	g.lastPlayer = player
}

func (g *Game) updateBoard(tile *Tile) {
	g.board.AddTileAt(tile)
}

// Winner returns the winning player or noPlayer if no winner.
func (g *Game) Winner() Player {
	return g.board.findRowFullWithSamePlayer()
}

// Tile represents a position on the board and knows how to compare/update itself.
type Tile struct {
	x      Row
	y      Column
	player Player
}

// NewTile creates a new tile.
func NewTile(x Row, y Column, player Player) *Tile {
	return &Tile{x: x, y: y, player: player}
}

// Player returns the tile's player.
func (t *Tile) Player() Player {
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
func (t *Tile) UpdatePlayer(newPlayer Player) {
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

func (b *Board) isTilePlayedAt(x Row, y Column) bool {
	return b.findTileAt(NewTile(x, y, noPlayer)).IsNotEmpty()
}

// AddTileAt sets the given tile's player at its position.
func (b *Board) AddTileAt(tile *Tile) {
	b.findTileAt(tile).UpdatePlayer(tile.Player())
}

func (b *Board) findRowFullWithSamePlayer() Player {
	if b.isRowFull(firstRow) && b.isRowFullWithSamePlayer(firstRow) {
		return b.playerAt(firstRow, firstColumn)
	}

	if b.isRowFull(secondRow) && b.isRowFullWithSamePlayer(secondRow) {
		return b.playerAt(secondRow, firstColumn)
	}

	if b.isRowFull(thirdRow) && b.isRowFullWithSamePlayer(thirdRow) {
		return b.playerAt(thirdRow, firstColumn)
	}

	return noPlayer
}

func (b *Board) findTileAt(tile *Tile) *Tile {
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(tile) {
			return t
		}
	}
	return nil
}

func (b *Board) hasSamePlayer(x Row, y Column, otherX Row, otherY Column) bool {
	return b.tileAt(x, y).HasSamePlayerAs(b.tileAt(otherX, otherY))
}

func (b *Board) playerAt(x Row, y Column) Player {
	return b.tileAt(x, y).Player()
}

func (b *Board) tileAt(x Row, y Column) *Tile {
	return b.findTileAt(NewTile(x, y, noPlayer))
}

func (b *Board) isRowFull(row Row) bool {
	return b.isTilePlayedAt(row, firstColumn) &&
		b.isTilePlayedAt(row, secondColumn) &&
		b.isTilePlayedAt(row, thirdColumn)
}

func (b *Board) isRowFullWithSamePlayer(row Row) bool {
	return b.hasSamePlayer(row, firstColumn, row, secondColumn) &&
		b.hasSamePlayer(row, secondColumn, row, thirdColumn)
}
