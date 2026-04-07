package anagrams

import "testing"

func TestAnagrams(t *testing.T) {
	t.Run("start at zero", func(t *testing.T) {
		if Anagrams != 0 {
			t.Errorf("expected 0, got %d", Anagrams)
		}
	})
}
