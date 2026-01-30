package tictactoe

import "testing"

func TestTicTacToe(t *testing.T) {
	t.Run("have a winner", func(t *testing.T) {
		if Winner != "a player" {
			t.Errorf("expected 'a player', got '%s'", Winner)
		}
	})
}
