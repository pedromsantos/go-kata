package romannumerals

import "testing"

func TestRomanNumeral(t *testing.T) {
	t.Run("be a roman numeral", func(t *testing.T) {
		if RomanNumeral != "roman numeral" {
			t.Errorf("expected 'roman numeral', got '%s'", RomanNumeral)
		}
	})
}
