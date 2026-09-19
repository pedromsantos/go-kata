package lsp

// Oven is the abstraction every cooking appliance in this kata should
// honour.
type Oven interface {
	Cook(food string) error
}
