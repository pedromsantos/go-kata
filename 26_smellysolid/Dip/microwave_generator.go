package dip

import "fmt"

// MicrowaveGenerator is a concrete low-level dependency.
type MicrowaveGenerator struct{}

// Generate produces microwaves.
func (MicrowaveGenerator) Generate() {
	fmt.Println("generating microwaves")
}
