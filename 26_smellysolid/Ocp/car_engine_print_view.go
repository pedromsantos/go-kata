package ocp

import "fmt"

// CarEnginePrintView renders a CarEngineViewModel as plain text.
type CarEnginePrintView struct {
	Text string
}

// FillWith renders viewModel into the view's Text field.
func (v *CarEnginePrintView) FillWith(viewModel CarEngineViewModel) {
	v.Text = fmt.Sprintf("RPM: %v, Temp: %v", viewModel.Rpm, viewModel.Temperature)
}
