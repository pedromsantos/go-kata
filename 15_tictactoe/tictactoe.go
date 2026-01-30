package tictactoe

// Player represents a player
type Player int

const (
	PlayerX Player = iota
	PlayerO
)

// Row represents a board row
type Row int

const (
	Top Row = iota
	Middle
	Bottom
)

// Column represents a board column
type Column int

const (
	Left Column = iota
	Center
	Right
)

// Cell represents a board position
type Cell struct {
	Row    Row
	Column Column
}

// Equals checks if two cells are equal
func (c Cell) Equals(other Cell) bool {
	return c.Row == other.Row && c.Column == other.Column
}

// Turn represents a player's move
type Turn struct {
	Cell   Cell
	Player Player
}

// Equals checks if two turns are equal
func (t Turn) Equals(other Turn) bool {
	return t.Player == other.Player && t.Cell.Equals(other.Cell)
}

// TicTacToe interface
type TicTacToe interface {
	Play(turn Turn)
}

// Output interface for game notifications
type Output interface {
	PrintPlay(x int, y int, player string)
	PrintWinner(player string)
	PrintError(errorMessage string)
}
