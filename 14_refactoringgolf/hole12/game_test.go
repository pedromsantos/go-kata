//nolint:all // This is intentionally "smelly" code for refactoring practice - do not fix linting issues
package hole12

import "testing"

func TestGame(t *testing.T) {
	t.Run("should not allow player O to play first", func(t *testing.T) {
		game := NewGame()
		err := game.Play("O", 0, 0)
		if err == nil {
			t.Error("expected error for O playing first")
		}
	})

	t.Run("should not allow player x to play twice in a row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		err := game.Play("X", 1, 0)
		if err == nil {
			t.Error("expected error for X playing twice")
		}
	})

	t.Run("should not allow a player to play in last played position", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		err := game.Play("O", 0, 0)
		if err == nil {
			t.Error("expected error for playing on last played position")
		}
	})

	t.Run("should not allow a player to play in any played position", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		err := game.Play("X", 0, 0)
		if err == nil {
			t.Error("expected error for playing on any played position")
		}
	})

	t.Run("should declare player X as winner if it plays three in top row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		game.Play("X", 0, 1)
		game.Play("O", 1, 1)
		game.Play("X", 0, 2)

		if winner := game.Winner(); winner != "X" {
			t.Errorf("expected X to win, got '%s'", winner)
		}
	})

	t.Run("should declare player O as winner if it plays three in top row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 1, 0)
		game.Play("O", 0, 0)
		game.Play("X", 1, 1)
		game.Play("O", 0, 1)
		game.Play("X", 2, 2)
		game.Play("O", 0, 2)

		if winner := game.Winner(); winner != "O" {
			t.Errorf("expected O to win, got '%s'", winner)
		}
	})

	t.Run("should declare player X as winner if it plays three in middle row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 1, 0)
		game.Play("O", 0, 0)
		game.Play("X", 1, 1)
		game.Play("O", 0, 1)
		game.Play("X", 1, 2)

		if winner := game.Winner(); winner != "X" {
			t.Errorf("expected X to win, got '%s'", winner)
		}
	})

	t.Run("should declare player O as winner if it plays three in middle row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 1, 0)
		game.Play("X", 2, 1)
		game.Play("O", 1, 1)
		game.Play("X", 2, 2)
		game.Play("O", 1, 2)

		if winner := game.Winner(); winner != "O" {
			t.Errorf("expected O to win, got '%s'", winner)
		}
	})

	t.Run("should declare player X as winner if it plays three in bottom row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 2, 0)
		game.Play("O", 0, 0)
		game.Play("X", 2, 1)
		game.Play("O", 0, 1)
		game.Play("X", 2, 2)

		if winner := game.Winner(); winner != "X" {
			t.Errorf("expected X to win, got '%s'", winner)
		}
	})

	t.Run("should declare player O as winner if it plays three in bottom row", func(t *testing.T) {
		game := NewGame()
		game.Play("X", 0, 0)
		game.Play("O", 2, 0)
		game.Play("X", 1, 1)
		game.Play("O", 2, 1)
		game.Play("X", 0, 1)
		game.Play("O", 2, 2)

		if winner := game.Winner(); winner != "O" {
			t.Errorf("expected O to win, got '%s'", winner)
		}
	})
}
