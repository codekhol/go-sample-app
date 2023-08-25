package notifier

type Notification struct {
	Destination string
	Content     string
}

type Notifier interface {
	Send(n Notification) error
}
