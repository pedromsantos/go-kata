package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("be zero", func(t *testing.T) {
		if Stack != 0 {
			t.Errorf("expected 0, got %d", Stack)
		}
	})
}
