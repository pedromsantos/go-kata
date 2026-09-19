package ocp

// CarEngineStatusReportController violates OCP: every new report format
// needs a new method on this controller (and a new concrete view struct) --
// the controller must be edited, not extended, to add a case.
type CarEngineStatusReportController struct {
	viewModel CarEngineViewModel
}

// NewCarEngineStatusReportController builds a controller for viewModel.
func NewCarEngineStatusReportController(viewModel CarEngineViewModel) *CarEngineStatusReportController {
	return &CarEngineStatusReportController{viewModel: viewModel}
}

// DisplayEngineStatusReport constructs a concrete web view directly instead
// of receiving one injected -- the DIP-shaped half of this same violation.
func (c *CarEngineStatusReportController) DisplayEngineStatusReport() *CarEngineWebView {
	webView := &CarEngineWebView{}
	webView.FillWith(c.viewModel)
	return webView
}

// PrintEngineStatusReport constructs a concrete print view directly instead
// of receiving one injected -- the DIP-shaped half of this same violation.
func (c *CarEngineStatusReportController) PrintEngineStatusReport() *CarEnginePrintView {
	printView := &CarEnginePrintView{}
	printView.FillWith(c.viewModel)
	return printView
}
