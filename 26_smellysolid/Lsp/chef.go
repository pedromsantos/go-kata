package lsp

// Chef's type-assertion here is the diagnostic signature of the LSP
// violation in Microwave: a caller that can't just trust Oven.Cook().
type Chef struct{}

// Cook cooks food in oven, special-casing Microwave because it can't
// honour Oven's Cook contract on its own.
func (Chef) Cook(oven Oven, food string) {
	if m, ok := oven.(*Microwave); ok {
		m.CookMicrowaving(food)
		return
	}
	_ = oven.Cook(food)
}
