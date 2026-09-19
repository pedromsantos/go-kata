package dip

// Kitchen (high-level policy) directly constructs a concrete MicrowaveOven
// (low-level detail) -- it can't work with any other kind of oven. DIP
// violation (also reads as an OCP violation, the same compound shape as
// the Ocp controller).
type Kitchen struct {
	oven *MicrowaveOven
}

// NewKitchen builds a Kitchen, constructing its own oven.
func NewKitchen() *Kitchen {
	return &Kitchen{oven: NewMicrowaveOven()}
}

// CookDinner cooks using the kitchen's oven.
func (k *Kitchen) CookDinner() {
	k.oven.Cook()
}
