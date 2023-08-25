package notifier

import (
	"fmt"
)

type smtpNotifier struct {
	addr     string
	sender   string
	password string
}

func New(addr, sender, password string) Notifier {
	return &smtpNotifier{
		addr:     addr,
		sender:   sender,
		password: password,
	}
}

func (smtp *smtpNotifier) Send(n Notification) error {
	// SMTP implementation goes here.
	// For simplicity we are only printing and email to standard output.
	fmt.Printf("TO: <%s>\nSUBJECT: %s\nBODY: %s\n", n.Destination, "Welcome!", n.Content)

	return nil
}
