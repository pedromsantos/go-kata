package primefactors

import "testing"

func TestPrimeFactors(t *testing.T) {
	t.Run("be one", func(t *testing.T) {
		if PrimeFactors != 1 {
			t.Errorf("expected 1, got %d", PrimeFactors)
		}
	})
}
