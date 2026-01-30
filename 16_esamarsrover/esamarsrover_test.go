package esamarsrover

import "testing"

func TestRover(t *testing.T) {
	t.Run("be at Mars plateau", func(t *testing.T) {
		if Rover != "Mars plateau" {
			t.Errorf("expected 'Mars plateau', got '%s'", Rover)
		}
	})
}
