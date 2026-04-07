package statscalculator

import "testing"

func TestStatsCalculator(t *testing.T) {
	t.Run("start at zero", func(t *testing.T) {
		if StatsCalculator != 0 {
			t.Errorf("expected 0, got %d", StatsCalculator)
		}
	})
}
