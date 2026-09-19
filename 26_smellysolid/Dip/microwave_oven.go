package dip

// MicrowaveOven news up a concrete MicrowaveGenerator itself instead of
// depending on an injected abstraction. DIP violation.
type MicrowaveOven struct {
	heater *MicrowaveGenerator
}

// NewMicrowaveOven builds a MicrowaveOven, constructing its own heater.
func NewMicrowaveOven() *MicrowaveOven {
	return &MicrowaveOven{heater: &MicrowaveGenerator{}}
}

// Cook heats using the oven's heater.
func (o *MicrowaveOven) Cook() {
	o.heater.Generate()
}
