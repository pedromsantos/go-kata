package tennis

// TennisGame interface for scoring tennis games.
type TennisGame interface {
	WonPoint(playerName string)
	GetScore() string
}

// TennisGame1 is the first implementation to refactor.
// This code is intentionally messy for refactoring practice.
type TennisGame1 struct {
	mScore1     int
	mScore2     int
	player1Name string
	player2Name string
}

// NewTennisGame1 creates a new tennis game.
func NewTennisGame1(player1Name, player2Name string) *TennisGame1 {
	return &TennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

// WonPoint records a point won by the named player.
func (g *TennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		g.mScore1++
	} else {
		g.mScore2++
	}
}

// GetScore returns the current score as a string.
func (g *TennisGame1) GetScore() string {
	score := ""
	tempScore := 0

	if g.mScore1 == g.mScore2 {
		switch g.mScore1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if g.mScore1 >= 4 || g.mScore2 >= 4 {
		minusResult := g.mScore1 - g.mScore2
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = g.mScore1
			} else {
				score += "-"
				tempScore = g.mScore2
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
