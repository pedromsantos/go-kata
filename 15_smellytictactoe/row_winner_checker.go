//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package smellytictactoe

// RowWinnerChecker checks whether a row on the board is won.
type RowWinnerChecker struct{}

// Check returns the winning symbol if a row is full with the same symbol, or " " otherwise.
func (c *RowWinnerChecker) Check(board *Board) string {
	//if the positions in first row are taken
	if board.TileAt(0, 0).Symbol != " " &&
		board.TileAt(0, 1).Symbol != " " &&
		board.TileAt(0, 2).Symbol != " " {
		//if first row is full with same symbol
		if board.TileAt(0, 0).Symbol == board.TileAt(0, 1).Symbol &&
			board.TileAt(0, 2).Symbol == board.TileAt(0, 1).Symbol {
			return board.TileAt(0, 0).Symbol
		}
	}

	//if the positions in second row are taken
	if board.TileAt(1, 0).Symbol != " " &&
		board.TileAt(1, 1).Symbol != " " &&
		board.TileAt(1, 2).Symbol != " " {
		//if middle row is full with same symbol
		if board.TileAt(1, 0).Symbol == board.TileAt(1, 1).Symbol &&
			board.TileAt(1, 2).Symbol == board.TileAt(1, 1).Symbol {
			return board.TileAt(1, 0).Symbol
		}
	}

	//if the positions in third row are taken
	if board.TileAt(2, 0).Symbol != " " &&
		board.TileAt(2, 1).Symbol != " " &&
		board.TileAt(2, 2).Symbol != " " {
		//if last row is full with same symbol
		if board.TileAt(2, 0).Symbol == board.TileAt(2, 1).Symbol &&
			board.TileAt(2, 2).Symbol == board.TileAt(2, 1).Symbol {
			return board.TileAt(2, 0).Symbol
		}
	}

	return " "
}
