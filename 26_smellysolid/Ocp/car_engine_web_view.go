package ocp

import "fmt"

// CarEngineWebView renders a CarEngineViewModel as HTML.
type CarEngineWebView struct {
	Html string
}

// FillWith renders viewModel into the view's Html field.
func (v *CarEngineWebView) FillWith(viewModel CarEngineViewModel) {
	v.Html = fmt.Sprintf("<div>%v rpm, %vC</div>", viewModel.Rpm, viewModel.Temperature)
}
