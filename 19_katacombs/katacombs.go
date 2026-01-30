package katacombs

// Game - implement this following TDD
type Game struct{}

// Output interface for game output
type Output interface {
	Print(message string)
}

// LocationRepository interface for locations
type LocationRepository interface {
	FindByTitle(title string) (interface{}, error)
}

// PlayerRepository interface for players
type PlayerRepository interface {
	Save(player interface{}) error
	FindByName(name string) (interface{}, error)
}
