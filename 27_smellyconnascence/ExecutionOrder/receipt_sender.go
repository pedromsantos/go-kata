package executionorder

import "fmt"

// ReceiptSender's Archive is only correct if SendToCustomer has already
// run, but nothing in the type or method signatures expresses that -- the
// caller has to know the right order from tribal knowledge, not from the
// code. Connascence of Execution Order.
type ReceiptSender struct {
	sent bool
}

// SendToCustomer emails the receipt to the customer.
func (s *ReceiptSender) SendToCustomer(receiptID string) {
	fmt.Printf("Emailing receipt %s to customer\n", receiptID)
	s.sent = true
}

// Archive archives the receipt, warning if it hasn't been sent yet.
func (s *ReceiptSender) Archive(receiptID string) {
	if !s.sent {
		fmt.Printf("Warning: archiving %s before it was sent\n", receiptID)
	}
	fmt.Printf("Archiving receipt %s\n", receiptID)
}
