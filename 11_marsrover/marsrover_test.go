package marsrover

import "testing"

func TestMarsRover(t *testing.T) {
	t.Run("execute returns input", func(t *testing.T) {
		rover := &Rover{}
		result := rover.Execute("test")
		if result != "test" {
			t.Errorf("expected 'test', got '%s'", result)
		}
	})
}
