package tennis

import "testing"

func TestTennisGame(t *testing.T) {
	testCases := []struct {
		p1Score  int
		p2Score  int
		expected string
	}{
		{0, 0, "Love-All"},
		{1, 1, "Fifteen-All"},
		{2, 2, "Thirty-All"},
		{3, 3, "Deuce"},
		{4, 4, "Deuce"},
		{1, 0, "Fifteen-Love"},
		{0, 1, "Love-Fifteen"},
		{2, 0, "Thirty-Love"},
		{0, 2, "Love-Thirty"},
		{3, 0, "Forty-Love"},
		{0, 3, "Love-Forty"},
		{4, 0, "Win for player1"},
		{0, 4, "Win for player2"},
		{2, 1, "Thirty-Fifteen"},
		{1, 2, "Fifteen-Thirty"},
		{3, 1, "Forty-Fifteen"},
		{1, 3, "Fifteen-Forty"},
		{4, 1, "Win for player1"},
		{1, 4, "Win for player2"},
		{3, 2, "Forty-Thirty"},
		{2, 3, "Thirty-Forty"},
		{4, 2, "Win for player1"},
		{2, 4, "Win for player2"},
		{4, 3, "Advantage player1"},
		{3, 4, "Advantage player2"},
		{5, 4, "Advantage player1"},
		{4, 5, "Advantage player2"},
		{6, 4, "Win for player1"},
		{4, 6, "Win for player2"},
	}

	for _, tc := range testCases {
		game := NewTennisGame1("player1", "player2")

		for i := 0; i < tc.p1Score; i++ {
			game.WonPoint("player1")
		}
		for i := 0; i < tc.p2Score; i++ {
			game.WonPoint("player2")
		}

		result := game.GetScore()
		if result != tc.expected {
			t.Errorf("Score %d-%d: expected '%s', got '%s'", tc.p1Score, tc.p2Score, tc.expected, result)
		}
	}
}
