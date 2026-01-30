package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("do something", func(t *testing.T) {
		if FizzBuzz != "something" {
			t.Errorf("expected 'something', got '%s'", FizzBuzz)
		}
	})
}
