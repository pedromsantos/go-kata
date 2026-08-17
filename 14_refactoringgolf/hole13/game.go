//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole13

import "errors"

// Row is a board row, 0 through 2.
type Row int

// Column is a board column, 0 through 2.
type Column int

// Player is a board symbol: noPlayer, "X" or "O".
type Player string

const (
	firstRow Row = 0
	thirdRow Row = 2

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
func (g *Game) Play(player Player, row Row, column Column) error {
	if err := g.validateFirstMove(player); err != nil {
		return err
	}
	if err := g.validatePlayer(player); err != nil {
		return err
	}
	coordinate := NewCoordinate(row, column)
	if err := g.validatePositionIsEmpty(coordinate); err != nil {
		return err
	}

	g.updateLastPlayer(player)
	g.updateBoard(NewTile(player, coordinate))
	return nil
}

func (g *Game) validateFirstMove(player Player) error {
	if g.lastPlayer == noPlayer && player == playerO {
		return errors.New("Invalid first player")
	}
	return nil
}

func (g *Game) validatePlayer(player Player) error {
	if player == g.lastPlayer {
		return errors.New("Invalid next player")
	}
	return nil
}

func (g *Game) validatePositionIsEmpty(coordinate Coordinate) error {
	if g.board.isTilePlayedAt(coordinate) {
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
	return g.board.findFullRowWithSamePlayerOrNoPlayer()
}

// Coordinate is a position on the board.
type Coordinate struct {
	x Row
	y Column
}

// NewCoordinate creates a new coordinate.
func NewCoordinate(x Row, y Column) Coordinate {
	return Coordinate{x: x, y: y}
}

// Equal reports whether two coordinates refer to the same position.
func (c Coordinate) Equal(other Coordinate) bool {
	return c.x == other.x && c.y == other.y
}

// Tile represents a played (or empty) position on the board.
type Tile struct {
	coordinate Coordinate
	player     Player
}

// NewTile creates a new tile.
func NewTile(player Player, coordinate Coordinate) *Tile {
	return &Tile{coordinate: coordinate, player: player}
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
	return t.coordinate.Equal(other.coordinate)
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
			b.plays = append(b.plays, NewTile(noPlayer, NewCoordinate(x, y)))
		}
	}
	return b
}

func (b *Board) isTilePlayedAt(coordinate Coordinate) bool {
	return b.findTileByCoordinate(NewTile(noPlayer, coordinate)).IsNotEmpty()
}

// AddTileAt sets the given tile's player at its position.
func (b *Board) AddTileAt(tile *Tile) {
	b.findTileByCoordinate(tile).UpdatePlayer(tile.Player())
}

func (b *Board) findFullRowWithSamePlayerOrNoPlayer() Player {
	for row := firstRow; row <= thirdRow; row++ {
		if b.isRowFull(row) && b.isRowFullWithSamePlayer(row) {
			return b.playerAt(NewCoordinate(row, firstColumn))
		}
	}

	return noPlayer
}

// findTileByCoordinate panics if the coordinate is off the board - this is a
// programmer-error invariant, not a recoverable validation failure.
func (b *Board) findTileByCoordinate(tile *Tile) *Tile {
	for _, t := range b.plays {
		if t.HasSameCoordinatesAs(tile) {
			return t
		}
	}

	panic("tile not found")
}

func (b *Board) playerAt(coordinate Coordinate) Player {
	return b.findTileByCoordinate(NewTile(noPlayer, coordinate)).Player()
}

func (b *Board) isRowFull(row Row) bool {
	return b.isTilePlayedAt(NewCoordinate(row, firstColumn)) &&
		b.isTilePlayedAt(NewCoordinate(row, secondColumn)) &&
		b.isTilePlayedAt(NewCoordinate(row, thirdColumn))
}

func (b *Board) isRowFullWithSamePlayer(row Row) bool {
	return b.hasSamePlayer(NewCoordinate(row, firstColumn), NewCoordinate(row, secondColumn)) &&
		b.hasSamePlayer(NewCoordinate(row, secondColumn), NewCoordinate(row, thirdColumn))
}

func (b *Board) hasSamePlayer(coordinate, otherCoordinate Coordinate) bool {
	return b.findTileByCoordinate(NewTile(noPlayer, coordinate)).HasSamePlayerAs(b.findTileByCoordinate(NewTile(noPlayer, otherCoordinate)))
}
