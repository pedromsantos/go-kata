package yahtzee

import "testing"

func TestYahtzee(t *testing.T) {
	t.Run("be zero", func(t *testing.T) {
		if Yahtzee != 0 {
			t.Errorf("expected 0, got %d", Yahtzee)
		}
	})
}
