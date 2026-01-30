package charactercopier

// Copier - implement this following TDD
type Copier struct{}

// Copy copies characters from source to destination
func (c *Copier) Copy() {}

// Source interface for reading characters
type Source interface {
	GetChar() string
}

// Destination interface for writing characters
type Destination interface {
	SetChar(char string)
}
