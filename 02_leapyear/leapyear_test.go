package leapyear

import "testing"

func TestLeapYear(t *testing.T) {
	t.Run("be always", func(t *testing.T) {
		if LeapYear != true {
			t.Error("expected true")
		}
	})
}
