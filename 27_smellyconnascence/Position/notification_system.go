package position

import "fmt"

// NotificationSystem's SendEmail takes three same-typed string parameters
// that carry meaning only through argument order -- swap recipient and
// sender and the call still compiles, but breaks silently. Connascence of
// Position.
type NotificationSystem struct{}

// SendEmail sends message from sender to recipient.
func (NotificationSystem) SendEmail(recipient, sender, message string) {
	fmt.Printf("From: %s\n", sender)
	fmt.Printf("To: %s\n", recipient)
	fmt.Printf("Message: %s\n", message)
}

// Demo exercises the positional ambiguity: nothing here stops recipient
// and sender from being silently swapped by a caller.
func Demo() {
	notificationSystem := NotificationSystem{}
	notificationSystem.SendEmail("recipient@email.com", "sender@email.com", "text")
}
