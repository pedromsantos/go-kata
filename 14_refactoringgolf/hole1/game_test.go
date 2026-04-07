//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole1

import "testing"

func TestGame(t *testing.T) {
	t.Run("X should win with first row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		game.Play("X", 0, 1)
		game.Play("O", 1, 1)
		game.Play("X", 0, 2)

		winner := game.Winner()
		if winner != "X" {
			t.Errorf("expected X to win, got '%s'", winner)
		}
	})

	t.Run("O should win with second row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		game.Play("X", 0, 1)
		game.Play("O", 1, 1)
		game.Play("X", 2, 2)
		game.Play("O", 1, 2)

		winner := game.Winner()
		if winner != "O" {
			t.Errorf("expected O to win, got '%s'", winner)
		}
	})

	t.Run("X always plays first", func(t *testing.T) {
		game := NewGame()
		err := game.Play("O", 0, 0)

		if err == nil {
			t.Error("expected error for O playing first")
		}
	})

	t.Run("players alternate", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		err := game.Play("X", 1, 0)

		if err == nil {
			t.Error("expected error for X playing twice")
		}
	})

	t.Run("cannot play on played position", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		err := game.Play("X", 0, 0)

		if err == nil {
			t.Error("expected error for playing on taken position")
		}
	})
}
