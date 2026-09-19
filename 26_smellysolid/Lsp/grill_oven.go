package lsp

import "fmt"

// GrillOven is a straightforward Oven implementation.
type GrillOven struct{}

// Cook grills food.
func (*GrillOven) Cook(food string) error {
	fmt.Printf("Grilling %s\n", food)
	return nil
}
