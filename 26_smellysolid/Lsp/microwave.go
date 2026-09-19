package lsp

import (
	"errors"
	"fmt"
)

// Microwave can't honour Oven's Cook contract, so it returns an error
// instead -- callers that only know about Oven get a broken promise.
// LSP violation.
type Microwave struct{}

// Cook always fails: Microwave does not support the Oven contract.
func (*Microwave) Cook(_ string) error {
	return errors.New("Microwave does not support Cook; use CookMicrowaving instead")
}

// CookMicrowaving is Microwave's real cooking method.
func (*Microwave) CookMicrowaving(food string) {
	fmt.Printf("Microwaving %s\n", food)
}
