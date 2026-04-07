package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	t.Run("start at zero", func(t *testing.T) {
		if Fibonacci != 0 {
			t.Errorf("expected 0, got %d", Fibonacci)
		}
	})
}
