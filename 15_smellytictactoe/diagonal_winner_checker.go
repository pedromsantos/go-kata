//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package smellytictactoe

// DiagonalWinnerChecker checks whether a diagonal on the board is won.
//
// Cross-file Duplicated Code / Shotgun Surgery kata fixture: this checker
// re-implements the exact same "are these three tiles taken and equal"
// pattern as RowWinnerChecker and ColumnWinnerChecker, independently,
// instead of sharing one extracted line-checking algorithm -- fixing a
// bug in the matching rule (e.g. wildcards, N-in-a-row) means
// remembering to edit all three files.
type DiagonalWinnerChecker struct{}

// Check returns the winning symbol if a diagonal is full with the same symbol, or " " otherwise.
func (c *DiagonalWinnerChecker) Check(board *Board) string {
	//if the positions in the top-left to bottom-right diagonal are taken
	if board.TileAt(0, 0).Symbol != " " &&
		board.TileAt(1, 1).Symbol != " " &&
		board.TileAt(2, 2).Symbol != " " {
		//if that diagonal is full with same symbol
		if board.TileAt(0, 0).Symbol == board.TileAt(1, 1).Symbol &&
			board.TileAt(2, 2).Symbol == board.TileAt(1, 1).Symbol {
			return board.TileAt(0, 0).Symbol
		}
	}

	//if the positions in the top-right to bottom-left diagonal are taken
	if board.TileAt(0, 2).Symbol != " " &&
		board.TileAt(1, 1).Symbol != " " &&
		board.TileAt(2, 0).Symbol != " " {
		//if that diagonal is full with same symbol
		if board.TileAt(0, 2).Symbol == board.TileAt(1, 1).Symbol &&
			board.TileAt(2, 0).Symbol == board.TileAt(1, 1).Symbol {
			return board.TileAt(0, 2).Symbol
		}
	}

	return " "
}
